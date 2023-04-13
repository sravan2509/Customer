CREATE TABLE `golang`.`customers` (
  `Name` VARCHAR(45) NULL,
  `Email` VARCHAR(45) NOT NULL,
  `PhoneNumber` VARCHAR(45) NULL,
  `Password` VARCHAR(100) NOT NULL,
  `Address` VARCHAR(45) NULL,
  PRIMARY KEY (`Email`),
  UNIQUE INDEX `Email_UNIQUE` (`Email` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;