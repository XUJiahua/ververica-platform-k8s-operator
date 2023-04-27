crds:
	helm upgrade -n evofinder \
  -i vp-k8s-operator-crds \
  ./charts/vp-k8s-operator-crds \
  --set .name=vp-k8s-operator-serving-cert \
  --set webhookService.name=vp-k8s-operator-webhook-service
controller:
	helm upgrade -n evofinder \
  -i vp-k8s-operator \
  ./charts/vp-k8s-operator \
  --set vvpEdition=community \
  --set vvpUrl=http://vvp-ververica-platform
