# go-mux-mvc-mysql-boilerplate
go-mux-mvc boilerplate with mysql database

## setup

* download this repo at your root project
* run the following script in your mysql workspace

    <code>
    CREATE DATABASE sample;

    DROP TABLE IF EXISTS `data`;
    CREATE TABLE `data` (
      `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
      `text` varchar(30) NOT NULL,
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
    </code>

## note

* router at ``controllers/controllers.go``

## env

* ``PORT``: ex ``8008``
