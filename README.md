# HOW USE IT

### Configure .env file
```
DB_User="root"
DB_Password=""
DB_Host="localhost"
DB_Port="3306"
DB_Name="qbcoreframework_39f588"
```
The migrated DB is inserted into `nf-skins` and `nf-outfits`.

```sql
CREATE TABLE IF NOT EXISTS `nf-outfits` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `citizenid` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `outfit_name` varchar(55) DEFAULT NULL,
  `share_code` varchar(55) DEFAULT uuid(),
  `model` int(11) DEFAULT NULL,
  `head_blend` text DEFAULT NULL,
  `face_feature` text DEFAULT NULL,
  `eye_color` int(11) DEFAULT NULL,
  `head_overlay` text DEFAULT NULL,
  `ped_component` text DEFAULT NULL,
  `ped_prop` text DEFAULT NULL,
  `tattoos` text DEFAULT NULL,
  `createdAt` timestamp NULL DEFAULT current_timestamp(),
  `updatedAt` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`) USING BTREE,
  KEY `FK_nf-skins_players` (`citizenid`) USING BTREE,
  CONSTRAINT `nf-outfits_ibfk_1` FOREIGN KEY (`citizenid`) REFERENCES `players` (`citizenid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci ROW_FORMAT=DYNAMIC;
```

```sql
CREATE TABLE IF NOT EXISTS `nf-skins` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `citizenid` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `model` int(11) DEFAULT NULL,
  `head_blend` text DEFAULT NULL,
  `face_feature` text DEFAULT NULL,
  `eye_color` int(11) DEFAULT NULL,
  `head_overlay` text DEFAULT NULL,
  `ped_component` text DEFAULT NULL,
  `ped_prop` text DEFAULT NULL,
  `tattoos` text DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_nf-skins_players` (`citizenid`),
  CONSTRAINT `FK_nf-skins_players` FOREIGN KEY (`citizenid`) REFERENCES `players` (`citizenid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci;
```
