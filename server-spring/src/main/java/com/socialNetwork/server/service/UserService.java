package com.socialNetwork.server.service;

import org.springframework.security.core.userdetails.UserDetailsService;

import com.socialNetwork.server.dto.UserLoginDTO;
import com.socialNetwork.server.dto.UserRegisterDTO;
import com.socialNetwork.server.entity.User;

public interface UserService extends UserDetailsService {
    boolean save(UserRegisterDTO userDTO);
    void save(User user);
    boolean login(UserLoginDTO userDTO);
}
