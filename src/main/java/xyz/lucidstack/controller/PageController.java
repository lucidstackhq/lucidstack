package xyz.lucidstack.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class PageController {

    @GetMapping("/")
    public String index() {
        return "index";
    }

    @GetMapping("/login")
    public String login() {
        return "login";
    }

    @GetMapping("/logout")
    public String logout() {
        return "logout";
    }

    @GetMapping("/join")
    public String join() {
        return "join";
    }

    @GetMapping("/home")
    public String home() {
        return "home";
    }

    @GetMapping("/account")
    public String account() {
        return "account";
    }

    @GetMapping("/organization")
    public String organization() {
        return "organization";
    }

    @GetMapping("/users")
    public String users() {
        return "users";
    }
}
