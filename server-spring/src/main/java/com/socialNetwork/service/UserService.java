package com.socialNetwork.service;

import org.springframework.security.core.userdetails.UserDetailsService;

import com.socialNetwork.dto.UserLoginDTO;
import com.socialNetwork.dto.UserRegisterDTO;
import com.socialNetwork.entity.User;

public interface UserService extends UserDetailsService {
    boolean save(UserRegisterDTO userDTO);
    void save(User user);
    boolean login(UserLoginDTO userDTO);
    User findUserByUsername(String username);
}
