CREATE TABLE `Authors`
(
    `ID`        int          NOT NULL,
    `LastName`  varchar(255) NOT NULL,
    `FirstName` varchar(255) NOT NULL,
    PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `Books`
(
    `ID`   int          NOT NULL,
    `Name` varchar(255) NOT NULL,
    PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `AuthorBook`
(
    `AuthorID` int NOT NULL,
    `BookID`   int NOT NULL,
    PRIMARY KEY (`AuthorID`, `BookID`),
    KEY        `BookID` (`BookID`),
    CONSTRAINT `AuthorBook_ibfk_1` FOREIGN KEY (`AuthorID`) REFERENCES `Authors` (`ID`),
    CONSTRAINT `AuthorBook_ibfk_2` FOREIGN KEY (`BookID`) REFERENCES `Books` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

