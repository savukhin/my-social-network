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
    this.auth.getUser().subscribe(
      user => {
        console.log("AppComponent User changed in app");
        
        this.user = user;
        this.cdref.detectChanges();
      }
    )

    setInterval(() => {
      let name = (Math.random() + 1).toString(36).substring(7);
      this.auth.setUser(new User(0, name));
    }, 1000)

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
