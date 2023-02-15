import { Component } from '@angular/core';
import { HttpService } from './http.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Quick Shop';
  items : any;
  showFiller = false;

  constructor(private httpService: HttpService) {}

  ngOnInit() {
    this.httpService.getWalmartItems("legos").subscribe(response => {
      console.log(response);
      this.items = response;
    })
  }
}
