import React from 'react';
import '../css/article.css'


const Article = (props) => {
    return (
        <div>
            <section className="s-content s-content--narrow s-content--no-padding-bottom">
                <article className="row format-standard">
                    <div className="s-content__header col-full">
                        <h1 className="s-content__header-title">
                            {props.title}
                        </h1>
                        <ul className="s-content__header-meta">
                            <li className="date">{props.createDate}</li>
                            <li className="cat">
                                {props.tags.map((value) => {
                                    return (
                                        <a href="#0">{value}</a>
                                    )
                                })}
                            </li>
                        </ul>
                    </div>

                    <div className="s-content__media col-full">
                        <div className="s-content__post-thumb">
                            <img src={props.image} alt=""/>
                        </div>
                    </div>

                    <div className="col-full s-content__main">
                        {props.content}
                    </div>
                </article>
            </section>
        </div>
    );
};

export default Article;