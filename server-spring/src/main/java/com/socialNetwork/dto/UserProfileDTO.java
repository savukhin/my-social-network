package com.socialNetwork.dto;

import java.sql.Date;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.Builder;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class UserProfileDTO {
    private Long id;
    private String username;
    private String name;
    private boolean isOnline;
    private String status;
    private Date birthDate;
    private String city;
    private String avatarURL;
}
