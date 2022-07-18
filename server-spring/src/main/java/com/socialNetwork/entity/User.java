package com.socialNetwork.entity;

import java.sql.Date;
import java.util.List;

import javax.persistence.*;

import org.springframework.lang.Nullable;

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
@Entity
@Table(name="users")
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String username;
    private String email;
    private String password;
    private String status;
    private String city;
    private Date birthDate;
    private boolean isOnline;

    @Nullable
    @OneToOne
    @JoinColumn(name = "photoId")
    private Content avatar;

    @OneToMany(cascade = CascadeType.ALL)
    private List<Content> photos;

    @OneToMany(cascade = CascadeType.ALL)
    private List<Content> posts;

    @OneToMany(cascade = CascadeType.ALL)
    private List<Like> likes;

    @Enumerated(EnumType.STRING)
    private Role role;
}
