create table achievement
(
    id              int auto_increment
        primary key,
    title           text not null,
    achievementText text not null,
    date            text not null,
    authorID        int  not null
);

create table user
(
    id        int auto_increment
        primary key,
    nickName  text not null,
    firstName text not null,
    lastName  text not null
);

create table comment
(
    id            int auto_increment
        primary key,
    userID        int  not null,
    achievementID int  not null,
    commentText   text not null
);

create table `like`
(
    id            int auto_increment
        primary key,
    userID        int not null,
    achievementID int null
);
