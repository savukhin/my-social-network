package com.socialNetwork.server.entity;

import java.util.List;

import javax.persistence.*;

@Entity
@Table(name="chats")
public class Chat {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String title;
    
    @OneToOne
    @JoinColumn(name = "photoId")
    private Content photo;

    @ManyToMany(cascade = CascadeType.ALL)
    private List<User> users;

    @OneToMany(cascade = CascadeType.ALL)
    private List<Content> messages;
}
