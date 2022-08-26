import { AfterViewInit, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { ActivatedRouteSnapshot, Router, RouterStateSnapshot } from '@angular/router';
import { catchError, map } from 'rxjs';
import { User, UserCompressed } from 'src/models/user';
import { AppearanceService } from './services/backend-api/appearance.service';
import { AuthService } from './services/backend-api/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements AfterViewInit, OnInit {
  private isLoaded = true;

  user?: UserCompressed;
  title = 'client-angular';
  blackout = false;

  constructor(private auth: AuthService, 
    private router: Router, 
    private cdref: ChangeDetectorRef,
    private appearance: AppearanceService) {

      const subscription = this.auth.getProfile(0)
      if (subscription != false) {
        
        subscription.subscribe(
          response => {
            if (response == false)
              return

            this.user = response;
            this.cdref.detectChanges();
          }
        )
      }

      this.appearance.blackout.subscribe(val => {
        this.blackout = val
      })

    // setTimeout(() => {
    //   let name = (Math.random() + 1).toString(36).substring(7);
    //   this.auth.setUser(new User(0, name));
    // }, 1000)

  }

  logoutClick() {
    this.auth.logout()
    this.router.navigateByUrl('login').then(() => {
      location.reload()
    })
  }

  blackoutClick() {
    this.appearance.blackout.next(false)
  }

  ngAfterViewInit(): void {
    let name = (Math.random() + 1).toString(36).substring(7);
      // this.auth.setUser(new User(0, name));
  }

  ngOnInit(): void {
    
  }
}
