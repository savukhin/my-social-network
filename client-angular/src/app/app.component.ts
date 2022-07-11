import { AfterViewInit, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { User } from 'src/models/user';
import { AuthService } from './services/backend-api/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [AuthService]
})
export class AppComponent implements AfterViewInit, OnInit {
  user?: User;
  title = 'client-angular';
  token = localStorage.getItem("token");

  constructor(private auth: AuthService, private cdref: ChangeDetectorRef) {
    this.auth.user$.subscribe(
      user => {
        console.log("User changed in app");
        
        this.user = user;
        this.cdref.detectChanges();
      }
    )
  }

  ngAfterViewInit(): void {
    
  }

  ngOnInit(): void {
    this.token = "asdf";
    
    if (this.token != null)
      this.auth.authWithToken(this.token)
        .subscribe((user) => {
          console.log(user);
          this.auth.setUser(user);
        })
  }
}
