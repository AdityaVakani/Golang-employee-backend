CREATE DATABASE  IF NOT EXISTS `go_backend01` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `go_backend01`;
-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: localhost    Database: go_backend01
-- ------------------------------------------------------
-- Server version	8.0.22

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
-- Table structure for table `department_lu`
--

DROP TABLE IF EXISTS `department_lu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `department_lu` (
  `depID` varchar(10) NOT NULL,
  `depName` varchar(100) NOT NULL,
  PRIMARY KEY (`depID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Look Up table for present departments and the id''s';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `department_lu`
--

LOCK TABLES `department_lu` WRITE;
/*!40000 ALTER TABLE `department_lu` DISABLE KEYS */;
INSERT INTO `department_lu` VALUES ('1','Department-1'),('2','Department-2'),('3','Department-3'),('4','Department-4');
/*!40000 ALTER TABLE `department_lu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `employee`
--

DROP TABLE IF EXISTS `employee`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `employee` (
  `empID` bigint NOT NULL AUTO_INCREMENT,
  `empName` varchar(100) NOT NULL,
  `empEmail` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `empPhone` varchar(45) NOT NULL,
  `empCity` varchar(255) NOT NULL,
  `empGender` enum('male','female','other') NOT NULL,
  `empDepartmentID` varchar(10) NOT NULL,
  `empHireDate` date NOT NULL,
  `empIsPermanent` tinyint(1) NOT NULL,
  PRIMARY KEY (`empID`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employee`
--

LOCK TABLES `employee` WRITE;
/*!40000 ALTER TABLE `employee` DISABLE KEYS */;
INSERT INTO `employee` VALUES (53,'rohan','rohan@gmail.com','98323678910','bangalore','female','2','2008-04-01',1),(54,'something','something@gmail.com','123456789','blore','male','4','1995-04-02',1),(58,'Mihir Veda','miveda91@yahoo.com','+919198623456','Mysore','male','2','2020-10-02',0),(67,'Aaa bbb ccc DD','xx@yy.zz','1234567890','aaa DD','male','3','2020-10-30',1);
/*!40000 ALTER TABLE `employee` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu` (
  `menuID` int NOT NULL AUTO_INCREMENT,
  `menuName` varchar(45) NOT NULL,
  `menuLink` varchar(45) DEFAULT '',
  `menuComponent` varchar(45) DEFAULT '',
  `menuVariant` varchar(45) DEFAULT '',
  `menuIcon` enum('IconLibraryBooks','IconPeople','IconBarChart','IconDashboard','IconShoppingCart','IconAccountBox','Default') DEFAULT 'Default',
  `parentID` int DEFAULT '0',
  PRIMARY KEY (`menuID`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (1,'Employees','/','','','IconPeople',0),(2,'Orders','/','','','IconShoppingCart',0),(7,'Dashboard','/','','','IconDashboard',0),(8,'Menu','/menu','','','IconLibraryBooks',0),(12,'ReportS','/reports','','','IconLibraryBooks',0),(14,'Report 2','/report2','','','IconDashboard',12),(15,'Report 55','/report5','','','IconLibraryBooks',12);
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mytable`
--

DROP TABLE IF EXISTS `mytable`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mytable` (
  `adult` varchar(126) NOT NULL,
  `belongs_to_collection` varchar(184) DEFAULT NULL,
  `budget` varchar(32) NOT NULL,
  `genres` varchar(264) NOT NULL,
  `homepage` varchar(242) DEFAULT NULL,
  `id` varchar(10) NOT NULL,
  `imdb_id` varchar(9) DEFAULT NULL,
  `original_language` varchar(5) DEFAULT NULL,
  `original_title` varchar(109) NOT NULL,
  `overview` varchar(1000) DEFAULT NULL,
  `popularity` varchar(21) DEFAULT NULL,
  `poster_path` varchar(35) DEFAULT NULL,
  `production_companies` varchar(1252) DEFAULT NULL,
  `production_countries` varchar(1039) DEFAULT NULL,
  `release_date` varchar(10) DEFAULT NULL,
  `revenue` int DEFAULT NULL,
  `runtime` decimal(6,1) DEFAULT NULL,
  `spoken_languages` varchar(765) DEFAULT NULL,
  `status` varchar(15) DEFAULT NULL,
  `tagline` varchar(297) DEFAULT NULL,
  `title` varchar(105) DEFAULT NULL,
  `video` varchar(5) DEFAULT NULL,
  `vote_average` decimal(4,1) DEFAULT NULL,
  `vote_count` int DEFAULT NULL,
  PRIMARY KEY (`adult`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mytable`
--

LOCK TABLES `mytable` WRITE;
/*!40000 ALTER TABLE `mytable` DISABLE KEYS */;
/*!40000 ALTER TABLE `mytable` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_details`
--

DROP TABLE IF EXISTS `order_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_details` (
  `productID` int NOT NULL AUTO_INCREMENT,
  `orderID` int NOT NULL,
  `product` varchar(45) NOT NULL,
  `quantity` int DEFAULT '1',
  `price` int NOT NULL,
  PRIMARY KEY (`productID`,`orderID`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_details`
--

LOCK TABLES `order_details` WRITE;
/*!40000 ALTER TABLE `order_details` DISABLE KEYS */;
INSERT INTO `order_details` VALUES (1,1,'3dprinter',1,19000),(3,1,'burger',1,300),(7,0,'4Dprinter',10,99000),(8,3,'prod1000',2,100),(9,3,'xxx',10,25),(13,4,'xxx',1,50),(14,4,'prod100',12,100);
/*!40000 ALTER TABLE `order_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `orderID` int NOT NULL AUTO_INCREMENT,
  `customer` varchar(100) DEFAULT '',
  `address` varchar(100) DEFAULT '',
  `city` varchar(45) DEFAULT '',
  `gender` enum('male','female','other') DEFAULT 'other',
  `orderDate` date DEFAULT '2000-01-01',
  `isDelivered` tinyint DEFAULT '0',
  PRIMARY KEY (`orderID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,'aditya','addr1','city1','male','2020-11-14',0),(3,'cust10','addr9','city11','other','2020-10-30',0),(4,'Mihir Vediya','Adarsh Palm Retreat','Mysore','male','2020-11-29',1),(6,'xx xx','zzzZ','aaa','male','2020-11-29',0);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-12-07 19:06:55
