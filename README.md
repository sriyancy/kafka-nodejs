# kafka-nodejs


1)kafka install
2).tgz file untar in c folder
3)rename file as kafka-open 
4)config->server.properties->edit->log.dirs=C:/kafka/kafka-logs
5)config->zookeeper.properties->edit->dataDir=C:/kafka/zookeeper-data
6)open cmd as admin go to c:/kafka
7)run zookeeper cmd:.\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties
8)run server cmd:.\bin\windows\kafka-server-start.bat .\config\server.properties
9)create topics
	cmd:c:\kafka>.\bin\windows\kafka-topics.bat --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic TestTopic 
---this is not cluster based

	cmd:list of topics:c:\kafka>.\bin\windows\kafka-topics.bat --list--zookeeper localhost:2181

10)add msgs  to topic::::
11)open cmd:.\bin\windows\kafka-console-producer.bat --broker-list localhost:9092 --topic TestTopic
12)open cmd to consumer creation for connsumeing data
	cmd:.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic TestTopic --from-beginning	


