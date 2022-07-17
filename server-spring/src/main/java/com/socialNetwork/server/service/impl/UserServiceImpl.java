package com.socialNetwork.server.service.impl;

import com.socialNetwork.server.service.UserService;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import com.socialNetwork.server.dto.UserLoginDTO;
import com.socialNetwork.server.dto.UserRegisterDTO;
import com.socialNetwork.server.entity.Role;
import com.socialNetwork.server.entity.User;
import com.socialNetwork.server.repository.UserRepository;

@Service
public class UserServiceImpl implements UserService {
    private final UserRepository userRepository;
    private final PasswordEncoder password_encoder;

    public UserServiceImpl( PasswordEncoder a, UserRepository b){
        this.password_encoder = a;
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
                .password(password_encoder.encode(userDTO.getPassword()))
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

        return password_encoder.encode(userDTO.getPassword()) == user.getPassword();
    }
}
