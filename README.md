[Automation with Kubernetes operators](https://jaxenter.de/kubernetes/automatisierung-kubernetes-operators-88189)

[kubernetes operator example](https://github.com/martin-helmich/kubernetes-operator-example)

err:
```bash
:~/go_client-go_workqueue/10-automation-with-kubernetes-operators/practise/src/github.com/ycliu912/helloworld-operator$ kubectl logs helloworld-operator-66458657d7-b5w65
{"level":"info","ts":1591865063.4795036,"logger":"cmd","msg":"Operator Version: 0.0.1"}
{"level":"info","ts":1591865063.4795759,"logger":"cmd","msg":"Go Version: go1.14.2"}
{"level":"info","ts":1591865063.479599,"logger":"cmd","msg":"Go OS/Arch: linux/amd64"}
{"level":"info","ts":1591865063.4796329,"logger":"cmd","msg":"Version of operator-sdk: v0.18.0+git"}
{"level":"info","ts":1591865063.4799178,"logger":"leader","msg":"Trying to become the leader."}
{"level":"info","ts":1591865064.0089152,"logger":"leader","msg":"Found existing lock with my name. I was likely restarted."}
{"level":"info","ts":1591865064.0089824,"logger":"leader","msg":"Continuing as the leader."}
{"level":"info","ts":1591865064.4896352,"logger":"controller-runtime.metrics","msg":"metrics server is starting to listen","addr":"0.0.0.0:8383"}
{"level":"info","ts":1591865064.514,"logger":"cmd","msg":"Registering Components."}
{"level":"error","ts":1591865064.5146556,"logger":"controller-runtime.eventhandler.EnqueueRequestForOwner","msg":"Could not get ObjectKinds for OwnerType","owner type":"<nil>","error":"expected pointer, but got invalid kind","stacktrace":"github.com/go-logr/zapr.(*zapLogger).Error\n\t/Users/liuyanchao/go/pkg/mod/github.com/go-logr/zapr@v0.1.1/zapr.go:128\nsigs.k8s.io/controller-runtime/pkg/handler.(*EnqueueRequestForOwner).parseOwnerTypeGroupKind\n\t/Users/liuyanchao/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.6.0/pkg/handler/enqueue_owner.go:97\nsigs.k8s.io/controller-runtime/pkg/handler.(*EnqueueRequestForOwner).InjectScheme\n\t/Users/liuyanchao/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.6.0/pkg/handler/enqueue_owner.go:179\nsigs.k8s.io/controller-runtime/pkg/runtime/inject.SchemeInto\n\t/Users/liuyanchao/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.6.0/pkg/runtime/inject/inject.go:98\nsigs.k8s.io/controller-runtime/pkg/manager.(*controllerManager).SetFields\n\t/Users/liuyanchao/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.6.0/pkg/manager/internal.go:254\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Watch\n\t/Users/liuyanchao/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.6.0/pkg/internal/controller/controller.go:127\ngithub.com/ycliu912/helloworld-operator/pkg/controller/helloworld.add\n\thelloworld-operator/pkg/controller/helloworld/helloworld_controller.go:60\ngithub.com/ycliu912/helloworld-operator/pkg/controller/helloworld.Add\n\thelloworld-operator/pkg/controller/helloworld/helloworld_controller.go:36\ngithub.com/ycliu912/helloworld-operator/pkg/controller.AddToManager\n\thelloworld-operator/pkg/controller/controller.go:13\nmain.main\n\thelloworld-operator/cmd/manager/main.go:126\nruntime.main\n\t/Users/liuyanchao/.g/go/src/runtime/proc.go:203"}
{"level":"error","ts":1591865064.5155802,"logger":"cmd","msg":"","error":"expected pointer, but got invalid kind","stacktrace":"github.com/go-logr/zapr.(*zapLogger).Error\n\t/Users/liuyanchao/go/pkg/mod/github.com/go-logr/zapr@v0.1.1/zapr.go:128\nmain.main\n\thelloworld-operator/cmd/manager/main.go:127\nruntime.main\n\t/Users/liuyanchao/.g/go/src/runtime/proc.go:203"}
```

local debug:
```bash
/private/var/folders/k8/t39hjy910bx2gh5k1ztk7fbc0000gn/T/___go_build_github_com_ycliu912_helloworld_operator_cmd_manager #gosetup
{"level":"info","ts":1591865356.984412,"logger":"cmd","msg":"Operator Version: 0.0.1"}
{"level":"info","ts":1591865356.984891,"logger":"cmd","msg":"Go Version: go1.14.2"}
{"level":"info","ts":1591865356.984915,"logger":"cmd","msg":"Go OS/Arch: darwin/amd64"}
{"level":"info","ts":1591865356.984944,"logger":"cmd","msg":"Version of operator-sdk: v0.18.0+git"}
{"level":"error","ts":1591865356.9849741,"logger":"cmd","msg":"Failed to get watch namespace","error":"WATCH_NAMESPACE must be set","stacktrace":"github.com/go-logr/zapr.(*zapLogger).Error\n\t/Users/liuyanchao/go_client-go_workqueue/10-automation-with-kubernetes-operators/practise/pkg/mod/github.com/go-logr/zapr@v0.1.1/zapr.go:128\nmain.main\n\t/Users/liuyanchao/go_client-go_workqueue/10-automation-with-kubernetes-operators/practise/src/github.com/ycliu912/helloworld-operator/cmd/manager/main.go:76\nruntime.main\n\t/Users/liuyanchao/.g/go/src/runtime/proc.go:203"}

Process finished with exit code 1

```

