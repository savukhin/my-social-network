import { Component, OnInit, ViewChild, ElementRef, AfterViewInit, ChangeDetectorRef } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from 'src/app/services/backend-api/auth.service';
import { User } from 'src/models/user';

@Component({
    selector: 'app-user-page',
    templateUrl: './user-page.component.html',
    styleUrls: ['./user-page.component.scss']
})
export class UserPageComponent implements AfterViewInit {
    user?: User;
    profile = new User(0,"Saveliy Karpukhin", true, "#notforwar", "30.08.2002", "Moscow")
    editStatus = false
    
    @ViewChild('pleaseDoIt') input: ElementRef<HTMLInputElement> = {} as ElementRef;

    showEditStatus() {
        setTimeout(() => {
            this.input.nativeElement.focus();
            this.input.nativeElement.value = this.profile.status;
        }, 0)
        this.editStatus = true;
    }

    hideEditStatus() {
        this.editStatus = false;
    }

    constructor(private route: ActivatedRoute, private router: Router, private auth: AuthService, private cdref: ChangeDetectorRef) {
        let id = this.route.snapshot.paramMap.get("id");
        if (id == null) {
            this.router.navigateByUrl('user/1')
            // this.router.navigateByUrl('user/' + this.auth.user.id, {skipLocationChange: true})
            return
        }

        console.log("subscription", id);
        this.auth.getProfile(+id).subscribe(
            user => {
                console.log("UserPageComponent: User changed in userpage");

                if (user) {
                    // this.profile = user;
                }
                this.cdref.detectChanges();
            }
        )
    }
    
    ngAfterViewInit(): void {
    }
}
