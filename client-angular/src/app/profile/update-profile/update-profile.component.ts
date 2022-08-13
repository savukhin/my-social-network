import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthService } from 'src/app/services/backend-api/auth.service';

@Component({
  selector: 'app-update-profile',
  templateUrl: './update-profile.component.html',
  styleUrls: ['./update-profile.component.scss']
})
export class UpdateProfileComponent implements OnInit {
  form: FormGroup

  constructor(private auth: AuthService, private fb: FormBuilder) { 
    this.form = this.fb.group({
      name: ['', Validators.required],
      birthDate: ['', Validators.required],
      city: ['', Validators.required],
    })
  }

  changeProfile() {
    if (!this.form.valid) 
      return
    
    const val = this.form.value

    let subscription = this.auth.changeProfile(val.name, val.birthDate, val.city)

    if (subscription == false)
      return

    subscription.subscribe((response) => {
        if (response == false) 
          console.log("Error");
        else
          console.log("ok");
      })
  }

  ngOnInit(): void {

  }

}
