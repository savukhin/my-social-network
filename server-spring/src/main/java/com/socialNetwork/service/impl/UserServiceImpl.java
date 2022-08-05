package com.socialNetwork.service.impl;

import com.socialNetwork.service.UserService;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import com.socialNetwork.dto.UserLoginDTO;
import com.socialNetwork.dto.UserRegisterDTO;
import com.socialNetwork.entity.Role;
import com.socialNetwork.entity.User;
import com.socialNetwork.repository.UserRepository;

@Service
public class UserServiceImpl implements UserService {
    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;

    // @Autowired
    // public void setUserRepository(UserRepository userRepository) {
    //     this.userRepository = userRepository;
    // }

    public UserServiceImpl(PasswordEncoder a, UserRepository b){
        this.passwordEncoder = a;
        this.userRepository = b;
    }

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        User user=userRepository.findFirstByUsername(username);
        if (user==null) {
            throw new UsernameNotFoundException("user is not found");
        }
        
        List<GrantedAuthority> roles=new ArrayList<>();
        roles.add(new SimpleGrantedAuthority(user.getRole().name()));

        return new org.springframework.security.core.userdetails.User(
                user.getUsername(),
                user.getPassword(),
                roles
        );
    }

    @Override
    public boolean save(UserRegisterDTO userDTO) {
        if (!Objects.equals(userDTO.getPassword(), userDTO.getPassword2())){
            throw new RuntimeException("passwords arent equal");
        }

        User user=User.builder()
                .username(userDTO.getUsername())
                .password(passwordEncoder.encode(userDTO.getPassword()))
                .email(userDTO.getEmail())
                .role(Role.USER)
                .build();

        userRepository.save(user);

        return true;
    }

    @Override
    public void save(User user) {
        userRepository.save(user);
    }

    @Override
    public boolean login(UserLoginDTO userDTO) {
        User user = userRepository.findByUsername(userDTO.getUsername());

        return passwordEncoder.encode(userDTO.getPassword()) == user.getPassword();
    }

    @Override
    public User findUserByUsername(String username) {
        return userRepository.findByUsername(username);
    }
}
