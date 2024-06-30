-- MariaDB dump 10.19-11.3.2-MariaDB, for osx10.19 (x86_64)
--
-- Host: localhost    Database: myblog
-- ------------------------------------------------------
-- Server version	8.1.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `blogs`
--
USE myblog;

DROP TABLE IF EXISTS `blogs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blogs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_by` varchar(255) NOT NULL,
  `created_at` int NOT NULL,
  `updated_at` int NOT NULL,
  `published_at` int NOT NULL,
  `status` tinyint NOT NULL,
  `title` varchar(255) NOT NULL,
  `author` varchar(255) NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `abstract` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Tags` varchar(255) NOT NULL,
  `audit_at` int NOT NULL,
  `audit_passed` binary(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_title` (`title`) USING BTREE,
  KEY `idx_author` (`author`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blogs`
--

LOCK TABLES `blogs` WRITE;
/*!40000 ALTER TABLE `blogs` DISABLE KEYS */;
INSERT INTO `blogs` VALUES
(1,'',1718759249,1719607206,1718782087,1,'New Title 3','testuser','New Content 34','New Abstact 3','{\"newone\":\"newone\"}',0,'0'),
(3,'',1718837688,1718837688,0,1,'from postman','testuser','from postman','from postman','{}',0,'0'),
(8,'',1718853751,1718853751,0,1,'from postman mq','testuser','from postmanmq','from postmanmq','{}',0,'0'),
(9,'testuser',1718900273,1718927477,0,1,'auth with token','testuser','auth with token','auth with token','{}',1718927477,'1'),
(11,'testuser',1719426889,1719607047,0,1,'How to Use md-editor-3','testuser','## 😲 md-editor-v3\n\nMarkdown Editor for Vue3, developed in jsx and typescript, support different themes、beautify content by prettier.\n\n### 🤖 Base\n\n**bold**, <u>underline</u>, _italic_, ~~line-through~~, superscript<sup>26</sup>, subscript<sub>1</sub>, `inline code`, [link](https://github.com/imzbf)\n\n> quote: I Have a Dream\n\n1. So even though we face the difficulties of today and tomorrow, I still have a dream.\n2. It is a dream deeply rooted in the American dream.\n3. I have a dream that one day this nation will rise up.\n\n- [ ] Friday\n- [ ] Saturday\n- [x] Sunday\n\n## 🤗 Code\n\n```vue\n<template>\n  <MdEditor v-model=\"text\" />\n</template>\n\n<script setup>\nimport { ref } from \'vue\';\nimport { MdEditor } from \'md-editor-v3\';\nimport \'md-editor-v3/lib/style.css\';\n\nconst text = ref(\'Hello Editor!\');\n</script>\n```\n\n## 🖨 Text\n\nThe Old Man and the Sea served to reinvigorate Hemingway\'s literary reputation and prompted a reexamination of his entire body of work.\n\n## 📈 Table\n\n| THead1          |      THead2       |           THead3 |\n| :-------------- | :---------------: | ---------------: |\n| text-align:left | text-align:center | text-align:right |\n\n## 📏 Formula\n\nInline: $x+y^{2x}$\n\n$$\n\\sqrt[3]{x}\n$$\n\n## 🧬 Diagram\n\n```mermaid\nflowchart TD\n  Start --> Stop\n```\n\n## 🪄 Alert\n\n!!! note Supported Types\n\nnote、abstract、info、tip、success、question、warning、failure、danger、bug、example、quote、hint、caution、error、attention\n\n!!!\n\n## ☘️ em...\n','Introduction to Styling in md-editor-3','{\"estew\":\"estew\",\"fest\":\"fest\",\"new\":\"new\",\"tet\":\"tet\"}',0,'0'),
(12,'testuser',1719427600,1719427611,0,1,'final test from ui','testuser','test content','final test from ui','{\"tag1\":\"tag1\"}',0,'0'),
(13,'ianzhang',1719608866,1719610140,1719610140,1,'Greetings from UI','ianzhang','# MyBlog: full stack go web app (CRUD)\n\n## Functionality\n\nmarkdown blog\n\n## Prototype\n\nUsers:\n\n- Guest viewers should be able to see the blogs without an account\n- creators should be able to publish blogs\n\nBasic workflow: publish, view blogs with/without accounts, comment\n\nPrototype:','Greetings from UI','{\"go\":\"go\",\"myblog\":\"myblog\"}',0,'0'),
(14,'ianzhang',1719611271,1719611273,1719611273,1,'my second blog','ianzhang','# This is my second blog','my second blog','{\"blog\":\"blog\",\"gin\":\"gin\",\"go\":\"go\"}',0,'0'),
(17,'ianzhang',1719681617,1719681622,1719681622,1,'Create a new blog','ianzhang','## use this editor to create a new blog','how to create ','{\"intro\":\"intro\",\"myblog\":\"myblog\"}',0,'0'),
(18,'ianzhang',1719682726,1719684853,1719684853,1,'how to use the blog app','ianzhang','## click the btns and find out, there must be some bugs hidding somewhere and i didn\'t find out. let me know if u find anything.this should be the last test from ui. im gonna deploy it soon. \n\n## this is a pretty cool editor btw','intro','{\"docker\":\"docker\",\"gin\":\"gin\",\"go\":\"go\",\"gorm\":\"gorm\",\"ioc\":\"ioc\",\"rabbitmq\":\"rabbitmq\",\"vuejs\":\"vuejs\"}',0,'0'),
(19,'ianzhang',1719683055,1719683055,1719683055,1,'last blog from ui, how to use the blog app','ianzhang','## last blog from ui\n\n\nclick the btns and find out. at this point i didn\'t find any bugs. but that doesn\'t mean there ain\'t no bugs. let me know if u find anything interesting','intro','{\"DDD\":\"DDD\",\"TDD\":\"TDD\",\"gin\":\"gin\",\"gorm\":\"gorm\",\"ioc\":\"ioc\",\"mysql\":\"mysql\",\"rabbitmq\":\"rabbitmq\",\"testing\":\"testing\",\"vuejs\":\"vuejs\"}',0,'0');
/*!40000 ALTER TABLE `blogs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_by` varchar(255) NOT NULL,
  `created_at` int NOT NULL,
  `blog_id` int NOT NULL,
  `content` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_blog_id` (`blog_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES
(1,'testuser',1719551930,1,'This is a test comment'),
(2,'ianzhang',1719592175,11,'good introduction to this editor'),
(3,'ianzhang',1719593495,11,'comment from postmand'),
(4,'ianzhang',1719600010,11,'another comment from postman'),
(5,'testuser',1719601077,11,'comment from ui'),
(6,'testuser',1719606112,9,'comment from ui'),
(7,'testuser',1719606175,9,'second comment from ui'),
(8,'ianzhang',1719681527,11,'good one'),
(9,'ianzhang',1719684630,19,'well actually i just ran into a few bugs, just fixed them');
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tokens`
--

DROP TABLE IF EXISTS `tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tokens` (
  `created_at` int NOT NULL,
  `updated_at` int NOT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `access_token` varchar(255) NOT NULL,
  `user_id` int NOT NULL,
  `refresh_token` varchar(255) NOT NULL,
  `access_token_expires_at` int NOT NULL,
  `refresh_token_expires_at` int NOT NULL,
  PRIMARY KEY (`access_token`) USING BTREE,
  UNIQUE KEY `idx_token` (`access_token`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tokens`
--

LOCK TABLES `tokens` WRITE;
/*!40000 ALTER TABLE `tokens` DISABLE KEYS */;
INSERT INTO `tokens` VALUES
(1719684789,1719684789,'ianzhang','cq04tdaclaah160e86vg',6,'cq04tdaclaah160e8700',7200,604800);
/*!40000 ALTER TABLE `tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` int NOT NULL,
  `updated_at` int NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` tinyint NOT NULL,
  `label` varchar(255) NOT NULL,
  `is_hashed` binary(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES
(4,1718499635,1718499635,'testuser','$2a$10$eT6x9NN4.pg2zcwL2S9/cOuI3KznwvmLwpNAVGJEHfWF.eRJVmvtO',0,'{}','1'),
(5,1718923839,1718923839,'audituser','$2a$10$eHOmZa1RZ6XNM6B9qgbrEO.3QzU/1W0WLJSsEGboUT2hX87e99tbS',2,'{}','1'),
(6,1719552391,1719552391,'ianzhang','$2a$10$AVG.lSA3Tqy5Fh6Jiw77v.pw/XiXs0s.2pxMSKEJtY8Jn8ptpovbq',0,'{}','1');
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

-- Dump completed on 2024-06-29 14:05:06
