# http2sns-vk

Designed to receive messages from S3 VK Cloud webhooks and forward messages to SNS.

### Run example

```sh
H2S_HTTP_LISTEN_ADDR=":5000" \
H2S_HTTP_ENDPOINT_TOPIC="/endpoint1:arn:aws:sns:eu-central-1:000000000000:SnsTopicName1,/endpoint2:arn:aws:sns:eu-central-1:000000000000:SnsTopicName2" \
H2S_HTTP_DEBUG=true \
H2S_SNS_API_ENDPOINT="http://some.sns.endpoint" \
H2S_REGION=eu-central-1 \
bin/http2sns
```