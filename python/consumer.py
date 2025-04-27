from consumer_setup import create_consumer
import json

# Replace with your Kafka credentials and topic
bootstrap_servers = ['Your Kafka bootstrap server url']
topic = 'Your Topic Name'  # Replace with your topic name
username = 'Your Username'  # Replace with your username
password = 'Your Password'  # Replace with your password

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
