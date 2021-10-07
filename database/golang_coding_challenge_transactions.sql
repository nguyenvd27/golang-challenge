CREATE DATABASE  IF NOT EXISTS `golang_coding_challenge` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `golang_coding_challenge`;
-- MySQL dump 10.13  Distrib 8.0.26, for macos11 (x86_64)
--
-- Host: localhost    Database: golang_coding_challenge
-- ------------------------------------------------------
-- Server version	8.0.26

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
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `amount` double DEFAULT NULL,
  `transaction_type` varchar(45) DEFAULT NULL,
  `account_id` int NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `account_id_idx` (`account_id`),
  CONSTRAINT `account_id` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (1,5000,'deposit',1,'2021-09-23 11:12:16','2021-09-23 11:12:16',NULL),(2,6000,'deposit',1,'2021-09-23 15:12:16','2021-09-23 15:12:16',NULL),(3,7000,'withdraw',2,'2021-09-23 15:40:16','2021-09-23 15:40:16',NULL),(4,8000,'withdraw',3,'2021-09-24 09:40:16','2021-09-24 09:40:16',NULL),(5,9000,'deposit',4,'2021-09-24 09:40:16','2021-09-24 09:40:16',NULL),(6,9877,'withdraw',4,'2021-09-24 11:10:07','2021-09-24 11:10:07',NULL),(7,9877,'withdraw',4,'2021-09-24 11:11:24','2021-09-24 11:11:24',NULL),(8,9877,'withdraw',4,'2021-09-24 11:26:07','2021-09-24 11:26:07',NULL),(9,1234.5,'deposit',4,'2021-09-24 11:29:20','2021-09-24 11:29:20',NULL),(10,222.54,'deposit',4,'2021-09-24 13:34:41','2021-09-24 13:34:41',NULL),(11,888.8,'deposit',3,'2021-09-24 13:35:33','2021-09-24 13:35:33',NULL),(12,777.89,'deposit',3,'2021-09-24 13:36:53','2021-09-24 13:36:53',NULL),(13,111.11,'withdraw',3,'2021-09-24 13:38:32','2021-09-24 13:38:32',NULL),(14,222.22,'withdraw',3,'2021-09-24 13:53:00','2021-09-24 13:53:00',NULL),(15,5555.67,'withdraw',1,'2021-09-27 14:35:34','2021-09-27 14:35:34',NULL),(16,45678.9,'withdraw',1,'2021-09-27 14:36:46','2021-09-27 14:36:46',NULL),(17,9999.98,'deposit',1,'2021-09-27 14:48:51','2021-09-27 14:48:51',NULL),(18,7899.89,'withdraw',1,'2021-09-30 14:48:51','2021-09-30 14:48:51',NULL),(19,12345.6,'deposit',1,'2021-10-01 14:48:51','2021-10-01 14:48:51',NULL);
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-10-01 17:07:13
