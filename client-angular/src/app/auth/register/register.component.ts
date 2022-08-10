import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/backend-api/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss'],
  providers: [AuthService, FormBuilder]
})
export class RegisterComponent implements OnInit {
  form: FormGroup;

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
  }

  register() {
    const val = this.form.value;

    if (val.email && val.password) {
        this.authService.register(val.login, val.email, val.password, val.password2)
            .subscribe(
                (response) => {
                    // console.log("User is logged in");
                    console.log(response);
                    // if (response.status == "ok")
                    //     this.router.navigateByUrl('/');
                }
            );
    }
}

}
