import { AfterViewInit, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { ActivatedRouteSnapshot, Router, RouterStateSnapshot } from '@angular/router';
import { User } from 'src/models/user';
import { AuthService } from './services/backend-api/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [AuthService]
})
export class AppComponent implements AfterViewInit, OnInit {
  private isLoaded = true;

  user?: User;
  title = 'client-angular';

  constructor(private auth: AuthService, 
    private router: Router, 
    private cdref: ChangeDetectorRef) {

      const subscription = this.auth.getProfile(0)
      if (subscription == false)
          return

      subscription.subscribe(
      response => {
        console.log("AppComponent User changed in app");
        if (response == false)
          return

        this.user = response;
        this.cdref.detectChanges();
      }
    )

    // setTimeout(() => {
    //   let name = (Math.random() + 1).toString(36).substring(7);
    //   this.auth.setUser(new User(0, name));
    // }, 1000)

  }

  ngAfterViewInit(): void {
    let name = (Math.random() + 1).toString(36).substring(7);
      // this.auth.setUser(new User(0, name));
  }

  ngOnInit(): void {
    
  }
}
