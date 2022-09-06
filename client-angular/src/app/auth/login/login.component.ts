import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { catchError, map, of } from 'rxjs';
import { AuthService } from 'src/app/services/backend-api/auth.service';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
    providers: [AuthService, FormBuilder]
})
export class LoginComponent implements OnInit {
    form: FormGroup;
    error?: string;

    constructor(private authService: AuthService, 
            private fb:FormBuilder, 
            private router: Router) {

        this.form = this.fb.group({
            login: ['',Validators.required],
            password: ['',Validators.required]
        });
    }

    ngOnInit(): void {
        if (this.authService.isAuthenticated)
            this.router.navigate(['user'])
    }

    login() {
        const val = this.form.value;

        if (val.login && val.password) {
            this.authService.login(val.login, val.password)
                .pipe(
                    map(
                        (response) => {
                            if (response.status == 200) {
                                this.router.navigate(['/user', response.body?.user_id])
                                    .then(() => { 
                                        location.reload()
                                    }
                                );
                            }
                        }
                    ),
                    catchError((response) => {
                        this.error = response.error.message
                        return of([])
                    })
                ).subscribe();
        }
    }

}
