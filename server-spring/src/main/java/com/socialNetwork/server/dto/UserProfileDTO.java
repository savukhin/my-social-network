package com.socialNetwork.server.dto;

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
    private boolean isOnline;
    private String status;
    private String birthDate;
    private String city;
    private String avatarURL;
}
