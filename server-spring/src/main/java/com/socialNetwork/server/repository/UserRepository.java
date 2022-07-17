package com.socialNetwork.server.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.socialNetwork.server.entity.User;

public interface UserRepository extends JpaRepository<User, Long> {
    User findByUsername(String username);
    User findFirstByUsername(String username);
}
