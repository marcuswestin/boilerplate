PROJECTNAME
===========

1: Change project name
----------------------

Replace projectname, and PROJECTNAME in all files with your project name, e.g "foobarcat, FOOBARCAT"

2: Setup development
--------------------

Run setup script:

	bash dev/_setup/one-time-setup.sh

Close window, open new terminal. From now on, to enter and activate project space:

	projectname

3: Kubernetes
-------------

Nodejs example:
    https://kubernetes.io/docs/tutorials/stateless-application/hello-minikube/#create-your-nodejs-application
Example guestbook go kubernetes file:
    https://github.com/kubernetes/kubernetes/blob/master/examples/guestbook-go/Dockerfile
Kube-solo info:
    https://github.com/TheNewNormal/kube-solo-osx
    Tutorial: https://deis.com/blog/2016/run-kubernetes-on-a-mac-with-kube-solo/
Example Kubereneters apps & archs:
    https://kubernetes.io/docs/samples/

Run kube-solo:

    open https://github.com/TheNewNormal/kube-solo-osx

See k8s dashboard:

    open http://192.168.64.2:8080/api/v1/proxy/namespaces/kube-system/services/kubernetes-dashboard/#/workload?namespace=default

Deploy dev app

    ./apps/dev/test/deploy.sh
