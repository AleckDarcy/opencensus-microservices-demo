export JAVA_HOME=$(/usr/libexec/java_home -v 1.8)

services=(
    "adservice"
    "checkoutservice"
    "currencyservice"
    "emailservice"
    "frontend"
    "paymentservice"
    "productcatalogservice"
    "shippingservice"
)

for ((i=0;i<${#services[@]};i++)) do
  cd src/${services[i]};

  pwd
  ./build.sh
  echo "done"

  cd ../..
done;