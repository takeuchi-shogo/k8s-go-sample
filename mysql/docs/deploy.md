
# MySQLをkubernetesにデプロイするまで

操作するディレクトリに移動してから

1 サンプルとなる.ymlファイルをイメージからコピーする

`kubectl create deploy mysql --image=mysql:5.7 --dry-run=client -o yaml > deployment.yml`

2 ネームスペースを作成する

`kubectl create ns database`

3 作成できたか確認する

`kubectl get ns`

4

`kubectl get pod -n database`

5

`kubectl apply -f='deployment.yml' -n database`

6

`kubectl get deploy`

7

`kubectl logs <pod name> -n <name space>`

8

`kubectl exec -n <name space> -it <pod name> -- mysql -u <root> -p <password>`

9

`kubectl delete deployment <deployment name> -n <name space>`
