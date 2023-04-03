-- MySQL dump 10.13  Distrib 8.0.27, for macos11 (x86_64)
--
-- Host: localhost    Database: operation
-- ------------------------------------------------------
-- Server version	5.7.28

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
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v6` varchar(25) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v7` varchar(25) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (56,'p','test','/api/system/menu','DELETE','','','','',''),(54,'p','test','/api/system/menu','POST','','','','',''),(55,'p','test','/api/system/menu','PUT','','','','',''),(53,'p','test','/api/system/menus','GET','','','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_system_login_log`
--

DROP TABLE IF EXISTS `tb_system_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_system_login_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `phone` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户账号',
  `ip` char(32) NOT NULL COMMENT 'IP地址',
  `address` varchar(256) NOT NULL COMMENT '登陆地址',
  `browser` varchar(128) NOT NULL COMMENT '浏览器',
  `device` varchar(128) NOT NULL COMMENT '登录设备',
  `status` tinyint(1) NOT NULL COMMENT '登录状态',
  `code` int(11) NOT NULL COMMENT '错误码',
  `description` varchar(256) NOT NULL COMMENT '登录备注',
  `created_at` int(11) DEFAULT NULL COMMENT '登陆时间',
  PRIMARY KEY (`id`),
  KEY `created_at` (`created_at`),
  KEY `phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_system_login_log`
--

LOCK TABLES `tb_system_login_log` WRITE;
/*!40000 ALTER TABLE `tb_system_login_log` DISABLE KEYS */;
INSERT INTO `tb_system_login_log` VALUES (1,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,1000015,'验证码错误',1678622656),(2,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,1000015,'验证码错误',1678622665),(3,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678622920),(4,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678623056),(5,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678623160),(6,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,4000,'token数据异常失败',1678623277),(7,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,4000,'token数据异常失败',1678623383),(8,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,1000014,'登陆密码时效已过期',1678623490),(9,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,4000,'token数据异常失败',1678623498),(10,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,1000014,'登陆密码时效已过期',1678623566),(11,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,4000,'token数据异常失败',1678623582),(12,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,1000014,'登陆密码时效已过期',1678623612),(13,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,100004,'数据库操作失败',1678623780),(14,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678623852),(15,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678624652),(16,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678624811),(17,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678624845),(18,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678624910),(19,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678625166),(20,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678625189),(21,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678625286),(22,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678625340),(23,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678625376),(24,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',0,1000014,'登陆密码时效已过期',1678625433),(25,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678625436),(26,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678625512),(27,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678625607),(28,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678626998),(29,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678627015),(30,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678627163),(31,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678628613),(32,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678628728),(33,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678636046),(34,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678719686),(35,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678846223),(36,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678885380),(37,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678888703),(38,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678888735),(39,'18288888888','::1','本地登陆','Chrome','macOS 10.15.7',1,0,'',1678890684),(40,'admin','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100005,'未查询到指定数据',1678941765),(41,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1678944689),(42,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1678944701),(43,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1678952830),(44,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1679076059),(45,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1679112531),(46,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1679127625),(47,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1679127633),(48,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1679127644),(49,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1679127659),(50,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1679127677),(51,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1679127837),(52,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1679128073),(53,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1679128182),(54,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100008,'账号密码错误',1679128192),(55,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1679128204),(56,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1679233489),(57,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100004,'数据库操作失败',1680236989),(58,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100004,'数据库操作失败',1680237240),(59,'18288888888','4.2.2.2','地址查询失败','Chrome','macOS 10.15.7',0,100004,'数据库操作失败',1680237330),(60,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,1000014,'登陆密码时效已过期',1680237422),(61,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',0,100004,'数据库操作失败',1680237483),(62,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1680237606),(63,'18288888888','4.2.2.2','地址查询失败','Chrome','macOS 10.15.7',1,0,'',1680238564),(64,'18288888888','4.2.2.2','地址查询失败','Chrome','macOS 10.15.7',1,0,'',1680238651),(65,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1680238843),(66,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1680239578),(67,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1680239622),(68,'18288888888','4.2.2.2',' 美国新泽西州纽瓦克市Level3Communications','Chrome','macOS 10.15.7',1,0,'',1680359288);
/*!40000 ALTER TABLE `tb_system_login_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_system_menu`
--

DROP TABLE IF EXISTS `tb_system_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_system_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '菜单标题',
  `icon` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '图标',
  `path` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '路径',
  `name` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '菜单name',
  `type` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '菜单类型',
  `permission` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '指令',
  `method` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '接口请求方式',
  `component` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '组件地址',
  `redirect` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '重定向地址',
  `parent_id` int(11) NOT NULL COMMENT '父级菜单ID',
  `is_hidden` tinyint(1) DEFAULT '0' COMMENT '是否隐藏',
  `is_cache` tinyint(1) DEFAULT '0' COMMENT '是否缓存页面',
  `weight` int(11) DEFAULT '0' COMMENT '菜单权重',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_system_menu`
--

LOCK TABLES `tb_system_menu` WRITE;
/*!40000 ALTER TABLE `tb_system_menu` DISABLE KEYS */;
INSERT INTO `tb_system_menu` VALUES (1,'根菜单',NULL,'/',NULL,'R','root',NULL,NULL,NULL,0,0,0,0,'system',0,1234567,1234567),(2,'系统管理','settings','/system','System','M','','','Layout','/system/user',1,0,1,0,'方伟业',1,1676616568,1678953420),(4,'基本接口','apps','/baseApi','baseApi','M','baseApi','','','',1,1,0,1,'方伟业',1,1676684535,1676685106),(5,'系统管理基础接口','apps','/baseApi','baseApi','M','','','','',4,1,0,0,'方伟业',1,1676684709,1676684897),(6,'获取当前用户信息','','/api/system/user','','A','baseApi','GET','','',5,0,0,0,'方伟业',1,1676684954,1676685013),(7,'获取当前用户菜单','','/api/system/user/menus','','A','baseApi','GET','','',5,0,0,0,'方伟业',1,1676685004,1676685004),(8,'获取系统部门信息','','/api/system/teams','','A','baseApi','GET','','',5,0,0,0,'方伟业',1,1676685078,1676685078),(9,'菜单管理','menu','menu','systemMenu','M','','','system/menu/index','',2,0,1,0,'方伟业',1,1676685381,1676685381),(10,'查看菜单','','/api/system/menus','','A','system:menu:query','GET','','',9,0,0,0,'方伟业',1,1676685528,1676685528),(11,'新增菜单','','/api/system/menu','','A','system:menu:add','POST','','',9,0,0,0,'方伟业',1,1676685599,1676685599),(12,'修改菜单','','/api/system/menu','','A','system:menu:update','PUT','','',9,0,0,0,'方伟业',1,1676685632,1676685632),(13,'删除菜单','','/api/system/menu','','A','system:menu:delete','DELETE','','',9,0,0,0,'方伟业',1,1676685657,1676685657),(14,'部门管理','user-group','team','sysTeam','M','','','system/team/index','',2,0,0,0,'方伟业',1,1676686013,1676691441),(15,'新增部门','','/api/system/team','','A','system:team:add','POST','','',14,0,0,0,'方伟业',1,1676686055,1676686055),(16,'修改部门','','/api/system/team','','A','system:team:update','PUT','','',14,0,0,0,'方伟业',1,1676686086,1676686086),(17,'删除部门','','/api/system/team','','A','system:team:delete','DELETE','','',14,0,0,0,'方伟业',1,1676686120,1676686120),(18,'角色管理','safe','role','sysRole','M','','','system/role/index','',2,0,0,0,'方伟业',1,1676686294,1676691447),(19,'查看角色','','/api/system/roles','','A','system:role:query','GET','','',18,0,0,0,'方伟业',1,1676686334,1676686334),(20,'新增角色','','/api/system/role','','A','system:role:add','POST','','',18,0,0,0,'方伟业',1,1676686390,1676691345),(21,'修改角色','','/api/system/role','','A','system:role:update','PUT','','',18,0,0,0,'方伟业',1,1676686414,1676691354),(22,'删除角色','','/api/system/role','','A','system:role:delete','DELETE','','',18,0,0,0,'方伟业',1,1676686455,1676691361),(23,'用户管理','user','user','sysUser','M','','','system/user/index','',2,0,0,0,'方伟业',1,1676686506,1676687642),(24,'查看用户','','/api/system/users','','A','system:user:query','GET','','',23,0,0,0,'方伟业',1,1676686542,1676686542),(25,'新增用户','','/api/system/user','','A','system:user:add','POST','','',23,0,0,0,'方伟业',1,1676686578,1676686648),(26,'修改用户','','/api/system/user','','A','system:user:update','PUT','','',23,0,0,0,'方伟业',1,1676686603,1676686643),(27,'删除用户','','/api/system/user','','A','system:user:delete','DELETE','','',23,0,0,0,'方伟业',1,1676686637,1676686637),(28,'登陆日志','history','login_log','sysLoginLog','M','','','system/login_log/index','',2,0,0,0,'方伟业',1,1676686716,1676687647),(29,'查看登陆日志','','/api/system/login/log','','A','system:login:log:query','GET','','',28,0,0,0,'方伟业',1,1676686777,1676686777),(30,'修改角色菜单','','','','G','system:role:menu','','','',18,0,0,0,'方伟业',1,1676687115,1678955234),(31,'获取角色的菜单id','','/api/system/role/menu_ids','','A','system:role:menu:query','GET','','',30,0,0,0,'方伟业',1,1676687183,1676687183),(32,'修改角色菜单','','/api/system/role/menu','','A','system:role:menu:update','PUT','','',30,0,0,0,'方伟业',1,1676687240,1676687240),(33,'数据展板','dashboard','/dashboard','Dashboard','M','','','Layout','/dashboard/workplace',1,0,0,1,'方伟业',1,1676687974,1678889506),(34,'系统数据','dashboard','workplace','Workplace','M','','','dashboard/workplace/index','',33,0,1,0,'方伟业',1,1676688043,1678889559),(35,'刷新用户token','','/api/system/token/refresh','','A','baseApi','POST','','',5,0,0,0,'',1,1678846301,1678846301);
/*!40000 ALTER TABLE `tb_system_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_system_role`
--

DROP TABLE IF EXISTS `tb_system_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_system_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '父角色id',
  `name` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '角色名称',
  `keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '角色关键字',
  `status` tinyint(1) NOT NULL COMMENT '角色状态',
  `weight` int(11) DEFAULT '0' COMMENT '角色权重',
  `description` varchar(300) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '角色备注',
  `data_scope` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '数据权限',
  `team_ids` text COLLATE utf8_unicode_ci COMMENT '自定义权限部门id',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_system_role`
--

LOCK TABLES `tb_system_role` WRITE;
/*!40000 ALTER TABLE `tb_system_role` DISABLE KEYS */;
INSERT INTO `tb_system_role` VALUES (1,0,'超级管理员','superAdmin',1,1,'超级管理员','ALLTEAM',NULL,'system',0,1676619290,1676619290),(2,1,'ss','test',1,0,'s','CUSTOM','[1]','方伟业',1,1678954131,1679129692);
/*!40000 ALTER TABLE `tb_system_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_system_role_menu`
--

DROP TABLE IF EXISTS `tb_system_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_system_role_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `menu_id` int(11) NOT NULL COMMENT '菜单ID',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`),
  KEY `menu_id` (`menu_id`),
  CONSTRAINT `tb_system_role_menu_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `tb_system_role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tb_system_role_menu_ibfk_2` FOREIGN KEY (`menu_id`) REFERENCES `tb_system_menu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_system_role_menu`
--

LOCK TABLES `tb_system_role_menu` WRITE;
/*!40000 ALTER TABLE `tb_system_role_menu` DISABLE KEYS */;
INSERT INTO `tb_system_role_menu` VALUES (81,2,34,'方伟业',1,1679135042,1679135042),(82,2,33,'方伟业',1,1679135042,1679135042),(83,2,10,'方伟业',1,1679135042,1679135042),(84,2,11,'方伟业',1,1679135042,1679135042),(85,2,12,'方伟业',1,1679135042,1679135042),(86,2,13,'方伟业',1,1679135042,1679135042),(87,2,9,'方伟业',1,1679135042,1679135042),(88,2,1,'方伟业',1,1679135042,1679135042),(89,2,2,'方伟业',1,1679135042,1679135042);
/*!40000 ALTER TABLE `tb_system_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_system_team`
--

DROP TABLE IF EXISTS `tb_system_team`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_system_team` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '部门名称',
  `description` varchar(300) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '部门备注',
  `parent_id` int(11) NOT NULL COMMENT '上级ID',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_system_team`
--

LOCK TABLES `tb_system_team` WRITE;
/*!40000 ALTER TABLE `tb_system_team` DISABLE KEYS */;
INSERT INTO `tb_system_team` VALUES (1,'青岑云科技','青岑云科技',0,'system',0,1673682377,1673682377),(2,'测试部门','ss',1,'方伟业',1,1679132308,1679132308);
/*!40000 ALTER TABLE `tb_system_team` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_system_user`
--

DROP TABLE IF EXISTS `tb_system_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_system_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `team_id` int(11) NOT NULL COMMENT '部门ID',
  `nickname` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户昵称',
  `name` varchar(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户姓名',
  `phone` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户电话',
  `avatar` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户头像',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户邮箱',
  `sex` tinyint(1) NOT NULL COMMENT '用户性别',
  `password` varchar(300) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户密码',
  `status` tinyint(1) NOT NULL COMMENT '用户状态',
  `disable_desc` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '禁用原因',
  `last_login` int(11) DEFAULT NULL COMMENT '最后登陆时间',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  UNIQUE KEY `email` (`email`),
  KEY `role_id` (`role_id`),
  KEY `team_id` (`team_id`),
  CONSTRAINT `tb_system_user_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `tb_system_role` (`id`),
  CONSTRAINT `tb_system_user_ibfk_2` FOREIGN KEY (`team_id`) REFERENCES `tb_system_team` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_system_user`
--

LOCK TABLES `tb_system_user` WRITE;
/*!40000 ALTER TABLE `tb_system_user` DISABLE KEYS */;
INSERT INTO `tb_system_user` VALUES (1,1,1,'柠檬很酸','方伟业','18288888888','','128@qq.com',1,'$2a$10$PnFNf65Ll6aDP0H/fFSI8OFXhK3UzJypkIza0PQOy7ze.98DM5ASa',1,'',1680359287,'方伟业',1,1676626494,1680359287);
/*!40000 ALTER TABLE `tb_system_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'operation'
--

--
-- Dumping routines for database 'operation'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-03  9:42:52
