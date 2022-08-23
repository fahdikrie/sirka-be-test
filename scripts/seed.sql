-- CreateSchema
CREATE SCHEMA IF NOT EXISTS TestCase;
SET search_path TO TestCase;

-- CreateTable
CREATE TABLE IF NOT EXISTS Users (
	Userid VARCHAR(50) NOT NULL,
	Name VARCHAR(50) NOT NULL,

	PRIMARY KEY (Userid)
);

-- Seed
INSERT INTO Users VALUES
	('01', 'budi'),
	('02', 'nano');