import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { catchError, map, of } from 'rxjs';
import { AuthService } from 'src/app/services/backend-api/auth.service';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss'],
    providers: [AuthService, FormBuilder]
})
export class RegisterComponent implements OnInit {
    form: FormGroup;
    error?: string;

    constructor(private authService: AuthService, 
            private fb:FormBuilder, 
            private router: Router) {

        this.form = this.fb.group({
            login: ['',Validators.required],
            email: ['',Validators.required],
            password: ['',Validators.required],
            password2: ['',Validators.required]
        });
    }

    ngOnInit(): void {
        if (this.authService.isAuthenticated)
            this.router.navigate(['user'])
    }

    register() {
        const val = this.form.value;

        if (val.email && val.password) {
            this.authService.register(val.login, val.email, val.password, val.password2)
                .pipe(
                    map(
                        (response) => {
                            if (!response.body) {
                                this.error = "Unknown error";
                                return;
                            }

                            if (response.status == 200) {
                                this.router.navigate(['/user', response.body.user_id])
                                    .then(() => { 
                                        location.reload()
                                    }
                                );
                            } else {
                                this.error = response.body.message
                                console.log("error = ", this.error);
                                
                            }
                        }
                    ),
                    catchError((response) => {
                        this.error = response.error.message
                        return of([])
                    })
                ).subscribe()
                
        }
    }

}
