import { AfterViewInit, Component } from '@angular/core';

interface Answer {
  title: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements AfterViewInit {
  title = 'client-angular';

  constructor() {}

  ngAfterViewInit(): void {
  }
}
