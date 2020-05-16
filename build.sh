export JAVA_HOME=$(/usr/libexec/java_home -v 1.8)

services=(
    "adservice"
#    "cartservice"
    "checkoutservice"
    "currencyservice"
#    "emailservice"
    "frontend"
    "paymentservice"
    "productcatalogservice"
#    "recommendationservice"
    "shippingservice"
)

for ((i=0;i<${#services[@]};i++)) do
  cd src/${services[i]};

  pwd
  sudo chmod +x build.sh
  ./build.sh
  echo "done"

  cd ../..
done;