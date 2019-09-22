import React from 'react';
import '../css/footer.css'
import '../css/base.css'

const Footer = () => {
    return (
        <footer class="s-footer">

            <div class="s-footer__main">
                <div class="row">

                    <div class="col-two md-four mob-full s-footer__archives">

                        <h4>Archives</h4>

                        <ul class="s-footer__linklist">
                            <li><a href="#0">September 2019</a></li>
                            <li><a href="#0">August 2019</a></li>
                            <li><a href="#0">July 2019</a></li>
                            <li><a href="#0">June 2019</a></li>
                        </ul>

                    </div>

                    <div class="col-two md-four mob-full s-footer__social">

                        <h4>Social</h4>

                        <ul class="s-footer__linklist">
                            <li><a href="#0">Github</a></li>
                            <li><a href="#0">Google+</a></li>
                            <li><a href="#0">LinkedIn</a></li>
                            <li><a href="#0">Facebook</a></li>
                        </ul>

                    </div>

                    <div class="col-six md-four mob-full end s-footer__subscribe">

                        <h4>Our Blogs</h4>

                        <p>Sit vel delectus amet officiis repudiandae est voluptatem. Tempora maxime provident nisi et fuga et enim exercitationem ipsam. Culpa consequatur occaecati.</p>

                        <div class="subscribe-form">
                            <form id="mc-form" class="group" novalidate="true">

                                <input type="email" value="" name="EMAIL" class="email" id="mc-email" placeholder="Email Address" required="" />

                                <label for="mc-email" class="subscribe-message"></label>

                            </form>
                        </div>
                    </div>

                </div>
            </div>

            <div class="s-footer__bottom">
                <div class="row">
                    <div class="col-full">
                        <div class="s-footer__copyright">
                            <span>© Copyright Dev Blogs 2019</span>
                            <span>Site Template by <a href="https://colorlib.com/">Colorlib</a></span>
                        </div>

                        <div class="go-top">
                            <a class="smoothscroll" title="Back to Top" href="#top"></a>
                        </div>
                    </div>
                </div>
            </div>
        </footer>
    );
};

export default Footer;