package com.socialNetwork.server.controller;

import java.io.IOException;

import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.servlet.mvc.method.annotation.ViewNameMethodReturnValueHandler;

import com.socialNetwork.server.dto.UserLoginDTO;
import com.socialNetwork.server.dto.UserRegisterDTO;
import com.socialNetwork.server.service.UserService;

@Controller
@RequestMapping("/user")
public class UserController {
    @Autowired
    private UserService userService;

    @RequestMapping(value="/registration", method=RequestMethod.POST)
    public void registration(@RequestBody UserRegisterDTO user, HttpServletResponse response) throws IOException {
        if (!user.getPassword().equals(user.getPassword2())) {
            response.sendError(HttpServletResponse.SC_BAD_REQUEST);
            return;
        }

        if (!userService.save(user)) {
            response.sendError(HttpServletResponse.SC_BAD_REQUEST);
            return;
        }

        response.setStatus(HttpServletResponse.SC_OK);
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
