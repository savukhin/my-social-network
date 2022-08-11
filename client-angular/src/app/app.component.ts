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
  private isLoaded = false;

  user?: User;
  title = 'client-angular';

  constructor(private auth: AuthService, 
    private router: Router, 
    private cdref: ChangeDetectorRef) {

    // this.auth.getUser().subscribe(
    //   user => {
    //     console.log("AppComponent User changed in app");
        
    //     this.user = user;
    //     this.cdref.detectChanges();
    //   }
    // )

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
    let token = this.auth.getToken();
    let unathorized_pages = new Set(["user", "login", "register"])
    let url = window.location.pathname
    let is_need_redirect = !(unathorized_pages.has(url.split('/')[1]))

    if (token != null) {
      this.auth.authWithToken(token).subscribe((response) => {
        this.isLoaded = true
        this.cdref.detectChanges()

        if (response.status != 200 && is_need_redirect)
         this.router.navigate(['/login'], { queryParams: { returnUrl: url }});
      })
    } else {
      this.isLoaded = true
      this.cdref.detectChanges()
      if (is_need_redirect)
        this.router.navigate(['/login'], { queryParams: { returnUrl: url }});
    }
  }
}
