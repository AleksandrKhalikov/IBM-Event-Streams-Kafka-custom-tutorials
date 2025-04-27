from kafka import KafkaConsumer

def create_consumer(topic, bootstrap_servers, username, password):
    # Set up the Kafka consumer with the provided configuration
    consumer = KafkaConsumer(
        topic,
        bootstrap_servers=bootstrap_servers,
        security_protocol='SASL_SSL',
        sasl_mechanism='SCRAM-SHA-512',
        sasl_plain_username=username,
        sasl_plain_password=password,
        ssl_cafile='certs/es-cert.pem',  # Ensure the correct path is set for your CA certificate
        auto_offset_reset='earliest',
        enable_auto_commit=True,
        group_id='my-python-consumer-group',
        value_deserializer=lambda x: x.decode('utf-8')
    )
    return consumer
