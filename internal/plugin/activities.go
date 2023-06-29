package plugin

import (
	"fmt"
	"strconv"
	"strings"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/known/durationpb"
)

// genActivitiesInterface generates an Activities interface
func (svc *Service) genActivitiesInterface(f *g.File) {
	f.Comment("Activities describes available worker activites")
	f.Type().Id("Activities").InterfaceFunc(func(methods *g.Group) {
		for _, activity := range svc.activitiesOrdered {
			method := svc.methods[activity]
			methods.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			methods.Id(activity).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(method.Output.GoIdent.GoName)
					}
					returnVals.Error()
				})
		}
	})
}

// genActivitiesInterface generates a RegisterActivities public function
func (svc *Service) genRegisterActivities(f *g.File) {
	f.Comment("RegisterActivities registers activities with a worker")
	f.Func().Id("RegisterActivities").
		Params(
			g.Id("r").Qual(workerPkg, "Registry"),
			g.Id("activities").Id("Activities"),
		).
		BlockFunc(func(fn *g.Group) {
			for _, activity := range svc.activitiesOrdered {
				fn.Id(fmt.Sprintf("Register%sActivity", activity)).Call(
					g.Id("r"), g.Id("activities").Dot(activity),
				)
			}
		})
}

// genRegisterActivity generates a Register<Activity> public function
func (svc *Service) genRegisterActivity(f *g.File, activity string) {
	method := svc.methods[activity]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	f.Commentf("Register%sActivity registers a %s activity", activity, activity)
	f.Func().Id(fmt.Sprintf("Register%sActivity", activity)).
		Params(
			g.Id("r").Qual(workerPkg, "Registry"),
			g.Id("fn").Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual("context", "Context")
					if hasInput {
						args.Op("*").Id(method.Input.GoIdent.GoName)
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(method.Output.GoIdent.GoName)
					}
					returnVals.Error()
				}),
		).
		Block(
			g.Id("r").Dot("RegisterActivityWithOptions").Call(
				g.Id("fn"), g.Qual(activityPkg, "RegisterOptions").Block(
					g.Id("Name").Op(":").Id(fmt.Sprintf("%sActivityName", activity)).Op(","),
				),
			),
		)
}

// genActivityFuture generates a <Activity>Future struct
func (svc *Service) genActivityFuture(f *g.File, activity string) {
	future := fmt.Sprintf("%sFuture", activity)

	f.Commentf("%s describes a %s activity execution", future, activity)
	f.Type().Id(future).Struct(
		g.Id("Future").Qual(workflowPkg, "Future"),
	)
}

// genActivityFutureGetMethod generates a <Workflow>Future's Get method
func (svc *Service) genActivityFutureGetMethod(f *g.File, activity string) {
	method := svc.methods[activity]
	hasOutput := !isEmpty(method.Output)
	future := fmt.Sprintf("%sFuture", activity)

	f.Commentf("Get blocks on a %s execution, returning the response", activity)
	f.Func().
		Params(g.Id("f").Op("*").Id(future)).
		Id("Get").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Id(method.Output.GoIdent.GoName)
				fn.If(
					g.Err().Op(":=").Id("f").Dot("Future").Dot("Get").Call(
						g.Id("ctx"), g.Op("&").Id("resp"),
					),
					g.Err().Op("!=").Nil(),
				).Block(
					g.Return(g.Nil(), g.Err()),
				)
				fn.Return(g.Op("&").Id("resp"), g.Nil())
			} else {
				fn.Return(g.Id("f").Dot("Future").Dot("Get").Call(
					g.Id("ctx"), g.Nil(),
				))
			}
		})
}

// genActivityFutureSelectMethod generates a <Workflow>Future's Select method
func (svc *Service) genActivityFutureSelectMethod(f *g.File, activity string) {
	future := fmt.Sprintf("%sFuture", activity)

	f.Commentf("Select adds the %s completion to the selector, callback can be nil", activity)
	f.Func().
		Params(g.Id("f").Op("*").Id(future)).
		Id("Select").
		Params(
			g.Id("sel").Qual(workflowPkg, "Selector"),
			g.Id("fn").Func().Params(g.Op("*").Id(future)),
		).
		Params(
			g.Qual(workflowPkg, "Selector"),
		).
		Block(
			g.Return(
				g.Id("sel").Dot("AddFuture").Call(
					g.Id("f").Dot("Future"),
					g.Func().
						Params(g.Qual(workflowPkg, "Future")).
						Block(
							g.If(g.Id("fn").Op("!=").Nil()).Block(
								g.Id("fn").Call(g.Id("f")),
							),
						),
				),
			),
		)
}

func (svc *Service) genSyncActivityFunction(f *g.File, activity string, local bool) {
	method := svc.methods[activity]
	methodName, hasInput, hasOutput := svc.getActivityDetails(method, local)

	f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if local {
				svc.addLocalActivityArgs(args, hasInput, hasOutput, method)
			}
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			callArgs := []g.Code{g.Id("ctx")}
			if hasInput {
				callArgs = append(callArgs, g.Id("req"))
			}
			if local {
				callArgs = append(callArgs, g.Id("fn"))
			}

			fn.Id("future").Op(":=").Id("Async" + methodName).Call(callArgs...)

			fn.Return(
				g.Id("future").Dot("Get").Call(g.Id("ctx")),
			)
		})
}

func (svc *Service) genAsyncActivityFunction(f *g.File, activity string, local bool) {
	method := svc.methods[activity]
	methodName, hasInput, hasOutput := svc.getActivityDetails(method, local)
	methodName = "Async" + methodName

	opts := svc.activities[activity].GetDefaultOptions()
	f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			if local {
				svc.addLocalActivityArgs(args, hasInput, hasOutput, method)
			}
		}).
		Params(
			g.Op("*").Id(fmt.Sprintf("%sFuture", method.GoName)),
		).
		BlockFunc(func(fn *g.Group) {
			svc.initializeActivityOptions(fn, local)
			svc.addActivityDefaultRetryPolicy(fn, opts)
			svc.addActivityDefaultTimeouts(fn, opts, local)
			svc.addCtxWithActivityOptions(fn, local)
			if local {
				svc.setActivityFunctionForLocal(fn, activity)
			}
			svc.addActivityFutureBuilder(fn, method, local, hasInput)
		})
}

func (svc *Service) getActivityDetails(method *protogen.Method, local bool) (string, bool, bool) {
	methodName := method.GoName
	if local {
		methodName = fmt.Sprintf("%sLocal", methodName)
	}
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	return methodName, hasInput, hasOutput
}

func (svc *Service) addLocalActivityArgs(args *g.Group, hasInput bool, hasOutput bool, method *protogen.Method) {
	args.Id("fn").
		Func().
		ParamsFunc(func(fnargs *g.Group) {
			fnargs.Qual("context", "Context")
			if hasInput {
				fnargs.Op("*").Id(method.Input.GoIdent.GoName)
			}
		}).
		ParamsFunc(func(fnreturn *g.Group) {
			if hasOutput {
				fnreturn.Op("*").Id(method.Output.GoIdent.GoName)
			}
			fnreturn.Error()
		})
}

func (svc *Service) initializeActivityOptions(fn *g.Group, local bool) {
	optionsFn := "GetActivityOptions"
	if local {
		optionsFn = "GetLocalActivityOptions"
	}
	fn.Id("opts").Op(":=").Qual(workflowPkg, optionsFn).Call(
		g.Id("ctx"),
	)
}

func (svc *Service) addActivityDefaultRetryPolicy(fn *g.Group, opts *temporalv1.ActivityOptions_StartOptions) {
	policy := opts.GetRetryPolicy()
	if policy == nil {
		return
	}

	fn.If(g.Id("opts").Dot("RetryPolicy").Op("==").Nil()).Block(
		g.Id("opts").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").ValuesFunc(func(fields *g.Group) {
			if d := policy.GetInitialInterval(); d.IsValid() {
				fields.Id("InitialInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
			}
			if d := policy.GetMaxInterval(); d.IsValid() {
				fields.Id("MaximumInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
			}
			if n := policy.GetBackoffCoefficient(); n != 0 {
				fields.Id("BackoffCoefficient").Op(":").Lit(n)
			}
			if n := policy.GetMaxAttempts(); n != 0 {
				fields.Id("MaximumAttempts").Op(":").Lit(n)
			}
			if errs := policy.GetNonRetryableErrorTypes(); len(errs) > 0 {
				fields.Id("NonRetryableErrorTypes").Op(":").Lit(errs)
			}
		}),
	)
}

func (svc *Service) addActivityDefaultTimeouts(fn *g.Group, opts *temporalv1.ActivityOptions_StartOptions, local bool) {
	if local {
		svc.addActivityDefaultTimeoutsLocal(fn, opts)
		return
	}

	svc.addActivityDefaultTimeoutsNonLocal(fn, opts)
}

func (svc *Service) addActivityDefaultTimeoutsNonLocal(fn *g.Group, opts *temporalv1.ActivityOptions_StartOptions) {
	svc.setTimeoutIfValid(fn, opts.GetHeartbeatTimeout(), "HeartbeatTimeout")
	svc.setTimeoutIfValid(fn, opts.GetScheduleToCloseTimeout(), "ScheduleToCloseTimeout")
	svc.setTimeoutIfValid(fn, opts.GetScheduleToStartTimeout(), "ScheduleToStartTimeout")
	svc.setTimeoutIfValid(fn, opts.GetStartToCloseTimeout(), "StartToCloseTimeout")
}

func (svc *Service) addActivityDefaultTimeoutsLocal(fn *g.Group, opts *temporalv1.ActivityOptions_StartOptions) {
	svc.setTimeoutIfValid(fn, opts.GetScheduleToCloseTimeout(), "ScheduleToCloseTimeout")
	svc.setTimeoutIfValid(fn, opts.GetStartToCloseTimeout(), "StartToCloseTimeout")
}

func (svc *Service) setTimeoutIfValid(fn *g.Group, timeout *durationpb.Duration, field string) {
	if !timeout.IsValid() {
		return
	}

	fn.If(g.Id("opts").Dot(field).Op("==").Lit(0)).Block(
		g.Id("opts").Dot(field).Op("=").
			Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).
			Comment(timeout.AsDuration().String()),
	)
}

func (svc *Service) addCtxWithActivityOptions(fn *g.Group, local bool) {
	if local {
		fn.Id("ctx").Op("=").Qual(workflowPkg, "WithLocalActivityOptions").Call(
			g.Id("ctx"),
			g.Id("opts"),
		)
		return
	}

	fn.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(
		g.Id("ctx"),
		g.Id("opts"),
	)
}

func (svc *Service) setActivityFunctionForLocal(fn *g.Group, activity string) {
	fn.Var().Id("activity").Any()
	fn.If(g.Id("fn").Op("==").Nil()).
		Block(
			g.Id("activity").Op("=").Id(fmt.Sprintf("%sActivityName", activity)),
		).
		Else().
		Block(
			g.Id("activity").Op("=").Id("fn"),
		)
}

func (svc *Service) addActivityFutureBuilder(fn *g.Group, method *protogen.Method, local bool, hasInput bool) {
	fn.Return(
		g.Op("&").Id(fmt.Sprintf("%sFuture", method.GoName)).ValuesFunc(func(bl *g.Group) {
			future := bl.Id("Future").Op(":")
			if local {
				future.Qual(workflowPkg, "ExecuteLocalActivity").CallFunc(func(returnVals *g.Group) {
					returnVals.Id("ctx")
					returnVals.Id("activity")
					if hasInput {
						returnVals.Id("req")
					}
				}).Op(",")
			} else {
				future.Qual(workflowPkg, "ExecuteActivity").CallFunc(func(returnVals *g.Group) {
					returnVals.Id("ctx")
					returnVals.Id(fmt.Sprintf("%sActivityName", method.GoName))
					if hasInput {
						returnVals.Id("req")
					}
				}).Op(",")
			}
		}),
	)
}
