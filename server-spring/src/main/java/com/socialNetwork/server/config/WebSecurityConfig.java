package com.socialNetwork.server.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.web.util.matcher.AntPathRequestMatcher;

import com.socialNetwork.server.entity.Role;

@Configuration
@EnableWebSecurity
public class WebSecurityConfig extends WebSecurityConfigurerAdapter {
    @Bean
    public PasswordEncoder passwordEncoder()
    {
        return new BCryptPasswordEncoder();
    }

    @Override
    protected void configure(HttpSecurity httpSecurity) throws Exception {
        httpSecurity.authorizeRequests()
                .antMatchers("/user").permitAll()
                .antMatchers("/user/registration").permitAll()
                .antMatchers("/user/login").permitAll()
                .anyRequest().permitAll()
                .and()
                    .csrf().disable();
                // .csrf()
                //     .disable()
                // .authorizeRequests()
                //     //Доступ только для не зарегистрированных пользователей
                //     .antMatchers("/user/registration").not().fullyAuthenticated()
                //     //Доступ только для пользователей с ролью Администратор
                //     .antMatchers("/admin/**").hasRole("ADMIN")
                //     .antMatchers("/news").hasRole("USER")
                //     //Доступ разрешен всем пользователей
                //     .antMatchers("/", "/resources/**").permitAll()
                // //Все остальные страницы требуют аутентификации
                // .anyRequest().authenticated()
                // .and()
                //     //Настройка для входа в систему
                //     .formLogin()
                //     .loginPage("/login")
                //     //Перенарпавление на главную страницу после успешного входа
                //     .defaultSuccessUrl("/")
                //     .permitAll()
                // .and()
                //     .logout()
                //     .permitAll()
                //     .logoutSuccessUrl("/");
    }
}
