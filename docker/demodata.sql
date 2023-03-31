LOCK TABLES `Authors` WRITE;

INSERT INTO `Authors` VALUES (1,'Иван','Иванов'),(2,'Петр','Петров');

UNLOCK TABLES;


LOCK TABLES `Books` WRITE;

INSERT INTO `Books` VALUES (1,'КНИГА ИВАНОВА'),(2,'КНИГА ИВАНОВА 2'),(3,'КНИГА ИВАНОВА И ПЕТРОВА');

UNLOCK TABLES;


LOCK TABLES `AuthorBook` WRITE;

INSERT INTO `AuthorBook` VALUES (1,1),(1,2),(1,3),(2,3);

UNLOCK TABLES;