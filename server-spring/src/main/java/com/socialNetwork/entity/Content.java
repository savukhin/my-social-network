package com.socialNetwork.entity;

import java.sql.Timestamp;
import java.util.List;

import javax.persistence.*;

import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.UpdateTimestamp;

@Entity
@Table(name="contents")
public class Content {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @CreationTimestamp
    private Timestamp creationTime;
    @UpdateTimestamp
    private  Timestamp updated;
    
    private String filepath;

    @Enumerated(EnumType.STRING)
    private ContentType type;

    @OneToMany(cascade = CascadeType.ALL)
    private List<Content> photos;

    @OneToMany(cascade = CascadeType.ALL)
    private List<Content> messages;

    @OneToMany(cascade = CascadeType.ALL)
    private List<Like> likes;
}
