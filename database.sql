CREATE DATABASE cad;

USE cad;

CREATE TABLE `user_profile` (
  `AccountID` int PRIMARY KEY AUTO_INCREMENT,
  `Name` varchar(255),
  `ProfileImg` varchar(255)
);

CREATE TABLE `user_details` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `AccountID` int,
  `IsDiabetic` bool,
  `Age` int,
  `Weight` float,
  `Height` float,
  `Gender` varchar(255)
);

CREATE TABLE `user_logs` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `AccountID` int,
  `GlucoseLevel` float,
  `Timestamp` int,
  `Classification` varchar(255)
);

CREATE TABLE `active_user` (
  `AccountID` int
);

CREATE TABLE `debug_glocuse_analog` (
  `VoltageData` float,
  `Timestamp` int
);

CREATE TABLE `active_recommendations` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `AccountID` int,
  `Recommendations` text
);

CREATE TABLE `classify` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `MinAge` int,
  `MaxAge` int,
  `MinGlocuseLevel` int,
  `MaxGlocuseLevel` int,
  `Classification` varchar(255)
);

CREATE TABLE `active_highblood_recommendations` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `AccountID` int,
  `Recommendations` text
);

CREATE TABLE `age_classification_recommendations` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `MinAge` int,
  `MaxAge` int,
  `Classification` varchar(255),
  `Recommendations` text
);

CREATE TABLE `bmi_classification_recommendations` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `MinBMI` float,
  `MaxBMI` float,
  `Classification` varchar(255),
  `Recommendations` text
);

CREATE TABLE `highblood_pressure_recommendations` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `Scenario` varchar(255),
  `MinBloodPressure` int,
  `MaxBloodPressure` int,
  `Recommendations` text
);

INSERT INTO classify (MinAge, MaxAge, MinGlocuseLevel, MaxGlocuseLevel, Classification) VALUES 
(6, 12, 0, 89, 'Low'),
(6, 12, 90, 140, 'Normal'),
(6, 12, 141, 999, 'High'),
(13, 19, 0, 89, 'Low'),
(13, 19, 90, 140, 'Normal'),
(13, 19, 141, 999, 'High'),
(20, 99, 0, 69, 'Low'),
(20, 99, 70, 140, 'Normal'),
(20, 99, 141, 999, 'High');