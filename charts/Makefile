crds:
	helm upgrade -n evofinder \
  -i vp-k8s-operator-crds \
  ./vp-k8s-operator-crds \
  --set .name=vp-k8s-operator-serving-cert \
  --set webhookService.name=vp-k8s-operator-webhook-service
controller-prod:
	helm upgrade -n evofinder \
  -i vp-k8s-operator \
  ./vp-k8s-operator \
  --set vvpEdition=community \
  --set vvpUrl=http://ververica-platform-ververica-platform
controller:
	helm upgrade -n evofinder \
  -i vp-k8s-operator \
  ./vp-k8s-operator \
  --set vvpEdition=community \
  --set vvpUrl=http://vvp-ververica-platform
certmgr:# controller required!
	kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.11.0/cert-manager.yaml
