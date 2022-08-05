import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/backend-api/auth.service';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
    providers: [FormBuilder]
})
export class LoginComponent implements OnInit {
    form: FormGroup;

    constructor(private authService: AuthService, 
            private fb:FormBuilder, 
            private router: Router) {

        this.form = this.fb.group({
            username: ['',Validators.required],
            password: ['',Validators.required]
        });
    }

    ngOnInit(): void {

    }

    login() {
        const val = this.form.value;

        if (val.username && val.password) {
            this.authService.login(
                val.username, 
                val.password,
                () => {
                    console.log("User is logged in");
                    // this.router.navigateByUrl('/');
                }
            )
        }
    }

}
