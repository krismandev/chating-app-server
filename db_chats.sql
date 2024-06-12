-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: db_chats
-- ------------------------------------------------------
-- Server version	5.7.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `chatroom_members`
--

DROP TABLE IF EXISTS `chatroom_members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chatroom_members` (
  `chatroom_id` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `join_date` datetime NOT NULL,
  `status` tinyint(1) NOT NULL,
  PRIMARY KEY (`chatroom_id`,`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chatroom_members`
--

LOCK TABLES `chatroom_members` WRITE;
/*!40000 ALTER TABLE `chatroom_members` DISABLE KEYS */;
INSERT INTO `chatroom_members` VALUES ('123','krisman','2024-06-11 08:59:31',1),('1234','krisman','2024-06-11 09:34:07',1),('abcd','krisman','2024-06-11 09:46:50',0),('dbtqi7fn','anto','2024-06-11 16:51:07',1),('dbtqi7fn','budi','2024-06-11 16:51:07',1),('dbtqi7fn','krisman','2024-06-11 16:51:07',1);
/*!40000 ALTER TABLE `chatroom_members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chatrooms`
--

DROP TABLE IF EXISTS `chatrooms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chatrooms` (
  `id` varchar(50) NOT NULL,
  `chatroom_name` varchar(45) NOT NULL,
  `username` varchar(50) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `chatroom_idx` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chatrooms`
--

LOCK TABLES `chatrooms` WRITE;
/*!40000 ALTER TABLE `chatrooms` DISABLE KEYS */;
INSERT INTO `chatrooms` VALUES ('abcd','Test 1','1','2024-06-11 00:28:22','2024-06-11 00:45:20'),('dbtqi7fn','Chatroom Satu','krisman','2024-06-11 16:47:53','2024-06-11 16:47:53'),('El1wd6mv','Test 2','1','2024-06-11 00:40:12','2024-06-11 00:40:12'),('olx61m2x','Test 3','krisman','2024-06-11 00:53:38','2024-06-11 00:53:38');
/*!40000 ALTER TABLE `chatrooms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chats`
--

DROP TABLE IF EXISTS `chats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chats` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message` text NOT NULL,
  `from` varchar(50) NOT NULL,
  `to` varchar(50) NOT NULL,
  `is_chatroom` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=79 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chats`
--

LOCK TABLES `chats` WRITE;
/*!40000 ALTER TABLE `chats` DISABLE KEYS */;
INSERT INTO `chats` VALUES (1,'halooo','krisman','budi',0,'2024-06-09 23:18:24'),(2,'sippp','krisman','budi',0,'2024-06-09 23:19:12'),(3,'okeee','krisman','1',0,'2024-06-11 13:43:49'),(4,'sipp','krisman','1',0,'2024-06-11 13:43:50'),(5,'','krisman','123',0,'2024-06-11 14:03:35'),(6,'sipp','krisman','123',0,'2024-06-11 14:05:40'),(7,'sipp deeeeehh','krisman','123',0,'2024-06-11 14:05:44'),(8,'exit','krisman','123',0,'2024-06-11 14:05:52'),(9,'/close','krisman','123',0,'2024-06-11 14:06:05'),(10,'sip','krisman','budi',0,'2024-06-11 14:10:03'),(11,'sip','krisman','budi',0,'2024-06-11 14:10:03'),(12,'sippp deh','krisman','budi',0,'2024-06-11 14:10:33'),(13,'sippp deh','krisman','budi',0,'2024-06-11 14:10:33'),(14,'halo 1','krisman','budi',0,'2024-06-11 14:10:54'),(15,'halo 1','krisman','budi',0,'2024-06-11 14:10:54'),(16,'okee','krisman','budi',0,'2024-06-11 14:16:05'),(17,'sippp','krisman','budi',0,'2024-06-11 14:16:40'),(18,'iyaaaa kris','budi','krisman',0,'2024-06-11 14:16:48'),(19,'test','krisman','budi',0,'2024-06-11 14:18:01'),(20,'oke','krisman','budi',0,'2024-06-11 14:18:04'),(21,'iyaa nih','budi','krisman',0,'2024-06-11 14:18:19'),(22,'terus gimana','krisman','budi',0,'2024-06-11 14:18:29'),(23,'','budi','krisman',0,'2024-06-11 14:18:36'),(24,'iyaa haha','budi','krisman',0,'2024-06-11 14:19:09'),(25,'oh okeee','krisman','budi',0,'2024-06-11 14:19:22'),(26,'sippp deh','krisman','budi',0,'2024-06-11 14:19:54'),(27,'','budi','krisman',0,'2024-06-11 14:20:22'),(28,'','krisman','budi',0,'2024-06-11 14:26:26'),(29,'testingg','krisman','dbtqi7fn',0,'2024-06-11 17:00:03'),(30,'test','krisman','2',0,'2024-06-11 17:12:13'),(31,'','krisman','2',0,'2024-06-11 17:12:22'),(32,'halo','krisman','2',0,'2024-06-11 17:26:51'),(33,'testingg','krisman','halooo',0,'2024-06-11 17:28:11'),(34,'okeee','krisman','halooo',0,'2024-06-11 17:28:18'),(35,'sip','krisman','2',0,'2024-06-11 17:30:40'),(36,'','krisman','2',0,'2024-06-11 17:31:33'),(37,'oke','krisman','dbtqi7fn',0,'2024-06-11 17:32:01'),(38,'halo','krisman','dbtqi7fn',0,'2024-06-11 17:32:03'),(39,'ok','budi','dbtqi7fn',0,'2024-06-11 17:32:08'),(40,'sip','budi','dbtqi7fn',0,'2024-06-11 17:32:10'),(41,'sip','krisman','dbtqi7fn',0,'2024-06-11 17:32:12'),(42,'','krisman','dbtqi7fn',0,'2024-06-11 17:32:13'),(43,'','krisman','dbtqi7fn',0,'2024-06-11 17:32:13'),(44,'sipppppp','krisman','dbtqi7fn',0,'2024-06-11 17:32:18'),(45,'haloooow','budi','dbtqi7fn',0,'2024-06-11 17:32:38'),(46,'sip','krisman','dbtqi7fn',0,'2024-06-11 17:48:01'),(47,'haloo','budi','dbtqi7fn',0,'2024-06-11 17:48:05'),(48,'iyaa dari anto','anto','dbtqi7fn',0,'2024-06-11 17:52:27'),(49,'iya dari budi','budi','dbtqi7fn',0,'2024-06-11 17:52:35'),(50,'iya dari krisman','krisman','dbtqi7fn',0,'2024-06-11 17:52:45'),(51,'sip','krisman','dbtqi7fn',0,'2024-06-11 17:53:01'),(52,'halo','anto','dbtqi7fn',0,'2024-06-11 18:02:04'),(53,'halo dari krisman','krisman','dbtqi7fn',0,'2024-06-11 18:08:30'),(54,'halo dri anto','anto','dbtqi7fn',0,'2024-06-11 18:08:43'),(55,'halo juga dari budi','budi','dbtqi7fn',0,'2024-06-11 18:08:58'),(56,'sipp','dimas','3',0,'2024-06-11 18:14:23'),(57,'haii dari budi lagi','budi','dbtqi7fn',0,'2024-06-11 18:14:32'),(58,'','budi','dbtqi7fn',0,'2024-06-11 19:21:50'),(59,'','krisman','dbtqi7fn',0,'2024-06-11 20:27:13'),(60,'','krisman','dbtqi7fn',0,'2024-06-11 20:27:13'),(61,'hai krisman','budi','krisman',0,'2024-06-11 21:15:03'),(62,'hai juga budi','krisman','budi',0,'2024-06-11 21:15:08'),(63,'testingg','budi','ok',0,'2024-06-11 21:35:39'),(64,'hau','budi','ok',0,'2024-06-11 21:35:43'),(65,'iyaa?','krisman','budi',0,'2024-06-11 21:35:47'),(66,'heii','krisman','budi',0,'2024-06-11 21:35:54'),(67,'haii dari krisman','krisman','budi',0,'2024-06-11 21:36:35'),(68,'hai dri budi','budi','ok',0,'2024-06-11 21:37:01'),(69,'','budi','ok',0,'2024-06-11 21:38:04'),(70,'hai dri budi','budi','krisman',0,'2024-06-11 21:38:25'),(71,'iya','budi','krisman',0,'2024-06-11 21:38:31'),(72,'okee','krisman','budi',0,'2024-06-11 21:38:34'),(73,'','budi','krisman',0,'2024-06-11 21:38:38'),(74,'','krisman','budi',0,'2024-06-11 21:38:40'),(75,'hai','krisman','dbtqi7fn',0,'2024-06-11 21:39:26'),(76,'oke iya?','budi','dbtqi7fn',0,'2024-06-11 21:39:31'),(77,'','krisman','dbtqi7fn',0,'2024-06-11 21:39:35'),(78,'','budi','dbtqi7fn',0,'2024-06-11 21:39:47');
/*!40000 ALTER TABLE `chats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `full_name` varchar(45) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `username` varchar(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Krisman Pratama',NULL,NULL,'$2a$08$AhRHrRvN1yS99YlSVv2.SO3a9u59RHvKNBdbsFmZYhFUBEAJl2lbq','krisman'),(2,'Budi',NULL,NULL,'$2a$08$wIBepishV8XpASrguJBvWe01xh4vHqoFfv8wminWJPegKMPt5fWwq','budi'),(3,'Anto',NULL,NULL,'$2a$08$Hl.50EDt/Od/DG57L.evfumiz0wfrQ.q0X4LESEWUxmS.Jx2EmDmm','anto'),(4,'Dimas',NULL,NULL,'$2a$08$s3vwomnjUFgI7WleIGCgoehmRov15T//sp/J7IjYgQVAfeyeT3aDO','dimas');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-11 21:45:31
