package com.socialNetwork.controller;

import java.io.IOException;
import java.security.Principal;

import javax.servlet.http.HttpSession;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

import com.socialNetwork.dto.UserProfileDTO;
import com.socialNetwork.dto.UserRegisterDTO;
import com.socialNetwork.entity.User;
import com.socialNetwork.service.UserService;

@Controller
@RequestMapping("/user")
public class UserController {
    @Autowired
    private UserService userService;

    @RequestMapping(value="/registration", method=RequestMethod.POST, produces = {"application/json"})
    public ResponseEntity<String> registration(@RequestBody UserRegisterDTO user) throws IOException {
        if (!user.getPassword().equals(user.getPassword2())) {
            return new ResponseEntity<>("{\"message\":\"Password doesn't match!\"}", HttpStatus.CONFLICT);
        }

        if (!userService.save(user)) {
            return new ResponseEntity<>("{\"message\":\"User with this password already exists\"}", HttpStatus.EXPECTATION_FAILED);
        }

        return new ResponseEntity<>("{\"message\":\"OK\"}", HttpStatus.CREATED);
    }

    @RequestMapping(value="/login", method=RequestMethod.GET)
    public ResponseEntity<String> login(Principal principal) {
        return new ResponseEntity<>("{\"message\": \"Sucess\"}", HttpStatus.OK);
    }

    @RequestMapping(value="/profile/{username}", method=RequestMethod.GET, produces = {"application/json"})
    // @CrossOrigin(origins="*", maxAge=3600)
    @CrossOrigin(origins={"http://localhost:4200"}, allowedHeaders = {"X-Auth-Token"}, maxAge=3600)
    public UserProfileDTO profileGET(@PathVariable String username) {
        User user = userService.findUserByUsername(username);

        UserProfileDTO dto = UserProfileDTO.builder()
            .id(user.getId())
            .username(user.getUsername())
            .name(user.getName())
            .isOnline(user.isOnline())
            .status(user.getStatus())
            .birthDate(user.getBirthDate())
            .city(user.getCity())
            .avatarURL(user.getAvatar().getFilepath())
            .build();

        return dto;
      }

    @RequestMapping(value="/token", method=RequestMethod.GET)
    public ResponseEntity<String> token(HttpSession session) {
        return ResponseEntity 
            .status(HttpStatus.OK)
            .body("{\"token\": \"" + session.getId() + "\"}");
    }
}
