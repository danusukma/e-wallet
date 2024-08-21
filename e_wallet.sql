/*
SQLyog Community v13.1.9 (64 bit)
MySQL - 5.6.17 : Database - e_wallet
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`e_wallet` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `e_wallet`;

/*Table structure for table `customer` */

DROP TABLE IF EXISTS `customers`;

CREATE TABLE `customers` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserName` longtext NOT NULL,
  `Password` longtext NOT NULL,
  `FullName` longtext NOT NULL,
  `Balance` bigint(20) DEFAULT 0,
  `DateCreated` datetime(3) DEFAULT current_timestamp(3),
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;

INSERT INTO e_wallet.customers (UserName,Password,FullName,Balance,DateCreated) VALUES
	 ('User-A','$2a$14$p2FrvtkXTWYoWFSv1ksePervNJ2ExVKW5kWtd0lk8ffj5QKAEjGGW','FullName-A',300,'2024-06-03 06:14:52.265000000'),
	 ('User-B','$2a$14$pIY7untbPqYqsorI9dLKNevThq8RHs1735OG8nBqKFqRTkYTETcbe','FullName-B',100,'2024-06-03 07:40:48.853000000');

/*Table structure for table `wallettransaction` */

DROP TABLE IF EXISTS `wallet_transactions`;

CREATE TABLE `wallet_transactions` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `DateTransaction` datetime(3) NOT NULL,
  `TypeTransaction` tinyint(4) NOT NULL,
  `FromId` bigint(20) NOT NULL,
  `ToId` bigint(20) NOT NULL,
  `Amount` bigint(20) NOT NULL DEFAULT 0,
  `DateCreated` datetime(3) DEFAULT current_timestamp(3),
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

INSERT INTO e_wallet.wallet_transactions (DateTransaction,TypeTransaction,FromId,ToId,Amount,DateCreated) VALUES
	 ('2024-06-03 08:56:06.420000000',1,12,13,100,'2024-06-03 08:56:06.421000000'),
	 ('2024-06-03 10:15:05.639000000',0,12,12,100,'2024-06-03 10:15:05.642000000'),
	 ('2024-06-03 10:15:05.639000000',1,13,12,50,'2024-06-03 11:47:17.813000000');


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
