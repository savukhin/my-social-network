package com.socialNetwork.controller;

import java.io.IOException;

import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

import com.socialNetwork.dto.UserLoginDTO;
import com.socialNetwork.dto.UserRegisterDTO;
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

    @RequestMapping(value="/login", method=RequestMethod.POST)
    public void registration(@RequestBody UserLoginDTO user, HttpServletResponse response) throws IOException {
        if (!userService.login(user)) {
            response.sendError(HttpServletResponse.SC_BAD_REQUEST);
            return;
        }

        response.setStatus(HttpServletResponse.SC_OK);
    }
}
