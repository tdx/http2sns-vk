# http2sns-vk

Designed to receive messages from S3 VK Cloud webhooks and forward messages to SNS.

### Run example

```sh
HTTP_LISTEN_ADDR=":5000" \
HTTP_ENDPOINT_TOPIC="/endpoint1:arn:aws:sns:eu-central-1:000000000000:SnsTopicName1,/endpoint2:arn:aws:sns:eu-central-1:000000000000:SnsTopicName2" \
HTTP_DEBUG=true \
SNS_API_ENDPOINT="http://some.sns.endpoint" \
bin/http2sns
```