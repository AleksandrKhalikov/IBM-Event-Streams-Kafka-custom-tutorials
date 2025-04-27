from consumer_setup import create_consumer
import json

# Replace with your Kafka credentials and topic
bootstrap_servers = ['es-demo-kafka-bootstrap-tools.apps.67eb6a82dbbf3678605ecc92.ap1.techzone.ibm.com:443']
topic = 'DOOR.BADGEIN'
username = 'python-consumer-1'
password = 'FDSrACfYignZqQPTJApM842jYJ8UlzWR'

# Create the Kafka consumer
consumer = create_consumer(topic, bootstrap_servers, username, password)

print(f"Consuming from topic: {topic}")

try:
    for message in consumer:
        print(f"Received message: {json.dumps(json.loads(message.value), indent=2)}")
except KeyboardInterrupt:
    print("Stopping consumer...")
finally:
    consumer.close()
