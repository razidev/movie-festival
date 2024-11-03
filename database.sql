CREATE DATABASE  IF NOT EXISTS `movie-festival` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `movie-festival`;
-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: localhost    Database: movie-festival
-- ------------------------------------------------------
-- Server version	8.2.0

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
-- Table structure for table `genres`
--

DROP TABLE IF EXISTS `genres`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `genres` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `viewers` int unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `genres`
--

LOCK TABLES `genres` WRITE;
/*!40000 ALTER TABLE `genres` DISABLE KEYS */;
INSERT INTO `genres` VALUES (1,'horror',0,'2024-11-02 14:16:37','2024-11-02 14:16:37'),(2,'superhero',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(3,'action',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(4,'comedy',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(5,'drama',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(6,'sci-fi',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(7,'fantasy',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(8,'thriller',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(9,'romance',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(10,'mystery',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(11,'animation',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(12,'documentary',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(13,'musical',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(14,'adventure',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(15,'crime',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(16,'biography',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(17,'historical',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(18,'western',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(19,'war',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(20,'sports',0,'2024-11-02 14:19:40','2024-11-02 14:19:40'),(21,'family',0,'2024-11-02 14:19:40','2024-11-02 14:19:40');
/*!40000 ALTER TABLE `genres` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movies`
--

DROP TABLE IF EXISTS `movies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movies` (
  `id` int unsigned NOT NULL,
  `unique_id` varchar(40) NOT NULL,
  `title` varchar(50) NOT NULL,
  `description` text,
  `duration` int unsigned DEFAULT NULL,
  `artists` json NOT NULL,
  `genre_ids` json NOT NULL,
  `watch_url` varchar(100) NOT NULL,
  `viewers` int unsigned DEFAULT '0',
  `voters` int DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_id_UNIQUE` (`unique_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movies`
--

LOCK TABLES `movies` WRITE;
/*!40000 ALTER TABLE `movies` DISABLE KEYS */;
/*!40000 ALTER TABLE `movies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_votes`
--

DROP TABLE IF EXISTS `user_votes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_votes` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `movie_unique_id` varchar(40) NOT NULL,
  `user_unique_id` varchar(40) NOT NULL,
  `status` enum('voted','unvoted') NOT NULL DEFAULT 'unvoted',
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_votes`
--

LOCK TABLES `user_votes` WRITE;
/*!40000 ALTER TABLE `user_votes` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_votes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `unique_id` varchar(40) NOT NULL,
  `email` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_id_UNIQUE` (`unique_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
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

-- Dump completed on 2024-11-02 21:20:29
