package com.socialNetwork.entity;

import java.util.List;

import javax.persistence.*;

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
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
